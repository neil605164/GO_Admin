package global

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var configFile []byte
var Config *DatabaseConfig

func Start() {
	var err error
	configFile, err = ioutil.ReadFile("conf/dev.yaml")
	if err != nil {
		log.Fatalf("yml file get err %v", err)
	}

	// 塞值進入struct
	Config, err = GetConfigData()
	if err != nil {
		panic(err)
	}

}

func GetConfigData() (db *DatabaseConfig, err error) {
	err = yaml.Unmarshal(configFile, &db)
	return db, err
}
