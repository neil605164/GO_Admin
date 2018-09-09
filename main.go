package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// func init() {
// 	// 顯示日期 ｜ 顯示時間（台灣） ｜ 顯示檔案名稱
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// }

// func main() {
// 	r := route.SetupRouter()

// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }

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

func main() {
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	fmt.Println(connectionString)
	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
}
