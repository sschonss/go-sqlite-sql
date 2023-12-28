package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		fmt.Println("Erro ao obter tabelas:", err)
		return
	}
	defer rows.Close()

	outputFile, err := os.Create("output.sql")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de saída:", err)
		return
	}
	defer outputFile.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Println("Erro ao escanear tabela:", err)
			return
		}

		createTableStmt := fmt.Sprintf("SELECT sql FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
		tableRows, err := db.Query(createTableStmt)
		if err != nil {
			fmt.Println("Erro ao obter o CREATE TABLE:", err)
			return
		}
		defer tableRows.Close()

		for tableRows.Next() {
			var createStmt string
			if err := tableRows.Scan(&createStmt); err != nil {
				fmt.Println("Erro ao escanear CREATE TABLE:", err)
				return
			}

			createStmt = strings.ReplaceAll(createStmt, "time@timestamp", "`time@timestamp`")

			_, err := outputFile.WriteString(fmt.Sprintf("%s;\n", createStmt))
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo de saída:", err)
				return
			}
		}

		dataRows, err := db.Query(fmt.Sprintf("SELECT * FROM \"%s\";", tableName))
		if err != nil {
			fmt.Println("Erro ao obter os dados da tabela:", err)
			return
		}
		defer dataRows.Close()

		columns, err := dataRows.Columns()
		if err != nil {
			fmt.Println("Erro ao obter colunas da tabela:", err)
			return
		}

		for dataRows.Next() {
			values := make([]interface{}, len(columns))
			columnPointers := make([]interface{}, len(columns))
			for i := range columns {
				columnPointers[i] = &values[i]
			}

			if err := dataRows.Scan(columnPointers...); err != nil {
				fmt.Println("Erro ao escanear linha da tabela:", err)
				return
			}

			insertStmt := fmt.Sprintf("INSERT INTO \"%s\" (%s) VALUES (", tableName, strings.Join(columns, ", "))
			valStrings := make([]string, len(columns))

			for i := range columns {
				switch values[i].(type) {
				case nil:
					valStrings[i] = "NULL"
				case []byte:
					valStrings[i] = fmt.Sprintf("'%s'", values[i])
				default:
					valStrings[i] = fmt.Sprintf("'%v'", values[i])
				}
			}

			insertStmt += strings.Join(valStrings, ", ") + ");\n"
			_, err := outputFile.WriteString(insertStmt)
			if err != nil {
				fmt.Println("Erro ao escrever comando INSERT no arquivo de saída:", err)
				return
			}
		}
	}

	fmt.Println("Conversão concluída. Arquivo output.sql criado com sucesso.")
}
