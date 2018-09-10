package dbConn

import (
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

var configFile []byte

// const (
// 	host     = "localhost"
// 	database = "GoAdmin"
// 	user     = "root"
// 	password = "qwe1234"
// )

type DatabaseConfig struct {
	Database Dbconnect `yaml:"database"`
}

type Dbconnect struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("conf/dev.yaml")
	if err != nil {
		log.Fatalf("yml file get err %v", err)
	}

	// 塞值進入struct
	config, err := GetDBConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}

func GetDBConfig() (db *DatabaseConfig, err error) {
	err = yaml.Unmarshal(configFile, &db)
	return db, err
}

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
