package main

import (
	"GO_Admin/global"
	"GO_Admin/model"
	route "GO_Admin/router"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	// 顯示日期 ｜ 顯示時間（台灣） ｜ 顯示檔案名稱
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 載入環境設定
	global.Start()
	// 檢查DB Table 是否存在
	err := model.CheckTableIsExist()
	if err != nil {
		panic(err)
	}
	// 載入router設定
	route.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
