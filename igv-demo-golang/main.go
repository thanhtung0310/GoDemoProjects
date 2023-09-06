package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"igv-demo-golang/database"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var connectionError error

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("Error loading credentials: %v", envErr)
	}

	var (
		password = os.Getenv("MSSQL_DB_PASSWORD")
		user     = os.Getenv("MSSQL_DB_USER")
		port     = os.Getenv("MSSQL_DB_PORT")
		dbName   = os.Getenv("MSSQL_DB_DATABASE")
	)

	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, dbName)

	sqlObj, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}

	data := database.Database{
		SqlDb: sqlObj,
	}

	fmt.Println("-> Welcome to IGV Console App, built using Golang and Microsoft SQL Server")
	fmt.Println("-> Select a numeric option; \n [1] Get list course \n [2] Get course by id \n [3] Select list web user \n [4] Select list admin user")

	consoleReader := bufio.NewScanner(os.Stdin)
	consoleReader.Scan()
	userChoice := consoleReader.Text()

	switch userChoice {
	case "1":
		data.GetListCourse()
		break

	case "2":
		break

	case "3":
		break

	case "4":
		break

	default:
		break
	}
}
