package dbConn

import (
	_ "github.com/go-sql-driver/mysql"
)

// func DBConnect() {
// 	// Initialize connection string.
// 	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?allowNativePasswords=true", user, password, host, database)

// 	// Initialize connection object.
// 	db, err := sql.Open("mysql", connectionString)
// 	checkError(err)
// 	defer db.Close()

// 	err = db.Ping()
// 	checkError(err)
// 	fmt.Println("Successfully created connection to database.")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
