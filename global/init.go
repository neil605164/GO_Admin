package global

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var configFile []byte

// Config 讀取dev.yaml檔案
var Config *DevConfig

// Start 執行main.go的第一步驟，載入設定檔
func Start() {
	var err error
	configFile, err = ioutil.ReadFile("conf/dev.yaml")
	if err != nil {
		log.Fatalf("yml file get err %v", err)
	}

	// 塞值進入struct
	Config, err = getConfigData()
	if err != nil {
		panic(err)
	}

}

func getConfigData() (db *DevConfig, err error) {
	err = yaml.Unmarshal(configFile, &db)
	return db, err
}
