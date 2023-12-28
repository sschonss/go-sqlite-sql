# SQLITE to SQL
## CLI for converting SQLITE to SQL

### Description

This CLI was created to convert SQLITE to SQL. It was used GoLang and libraries: [go-sqlite3](github.com/mattn/go-sqlite3 v1.14.19).

---

### Dependencies

- [GoLang](https://golang.org/dl/) >= v1.16.5
- [go-sqlite3](github.com/mattn/go-sqlite3 v1.14.19)
- SQLITE file with the name `database.db` in the same folder as the executable

---
### Installation

How to install the golang: [GoLang](https://golang.org/dl/)

#### Linux

1. Open the terminal
2. Run the command to download golang

```bash
sudo apt-get update && sudo apt-get install golang-go
```
3. Verify the installation

```bash
go version
```

#### Windows

1. Download the installer from the [official website](https://golang.org/dl/)
2. Run the installer
3. Verify the installation

```bash
go version
```

---

### How to use

1. Download the project

2. Run the command below to install the dependencies

```bash
go get
```

3. Run the command below to build the project

```bash
go build -o sqlite2sql
```

4. Save the SQLITE file in the same folder as the executable with the name `database.db`

5. Run the command below to convert the SQLITE file to SQL

```bash
./sqlite2sql
```
6. The SQL file will be created in the same folder as the executable with the name `output.sql`

---

### Author

- [Luiz Schons](github.com/sschons)

---

### License

[MIT](https://choosealicense.com/licenses/mit/)

