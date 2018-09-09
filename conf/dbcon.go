package dbConn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "goAdmindb"
	database = "GoAdmin"
	user     = "root"
	password = "qwe1234"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func DBConnect() {
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
}
