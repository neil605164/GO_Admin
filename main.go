package main

import (
	route "GO_Admin/router"
	"log"

	conf "GO_Admin/conf"
)

func init() {
	// 顯示日期 ｜ 顯示時間（台灣） ｜ 顯示檔案名稱
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	conf.DBConnect()
	r := route.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
