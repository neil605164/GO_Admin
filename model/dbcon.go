package dbConn

import (
	"GO_Admin/global"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	USER     = global.Config.Database.User
	PASSWORD = global.Config.Database.Password
	HOST     = global.Config.Database.Host
	DATABASE = global.Config.Database.Database
)

func DBConnect() (db *sql.DB) {

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?allowNativePasswords=true", USER, PASSWORD, HOST, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
