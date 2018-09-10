package global

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var configFile []byte
var config DatabaseConfig

func Start() {
	var err error
	configFile, err = ioutil.ReadFile("conf/dev.yaml")
	if err != nil {
		log.Fatalf("yml file get err %v", err)
	}

	// 塞值進入struct
	config, err := GetConfigData()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}

func GetConfigData() (db *DatabaseConfig, err error) {
	err = yaml.Unmarshal(configFile, &db)
	return db, err
}
