package helper

import (
	"fmt"
	"log"
	"os"
	"time"
)

// WriteLog 寫Log
func WriteLog() {
	// 取當前時間
	start := time.Now()
	logTime := start.Format("2006-01-02 15:04:05 -07:00")

	// 取得執行時間
	t := time.Now()
	elapsed := t.Sub(start)

	// 組合Log訊息
	logString := fmt.Sprintf("[INFO] %v | %v  ",
		logTime,
		elapsed)
	fmt.Println(logString)
	// 設定檔案位置
	fileName := "guava_access.log"
	filePath := "/home/hsieh/log/"

	// 建制資料夾 + 檔案
	if _, err := os.Stat(filePath); err != nil {
		if err := os.MkdirAll(filePath, 0777); err != nil {
			log.Printf("❌ WriteLog: 建立資料夾錯誤 [%v] ❌ \n----> %s\n", err)
			return
		}
	}

	_, err := os.Create(filePath + fileName)
	if err != nil {
		log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n----> %s\n", err)
		return
	}

	// 開啟檔案
	outputf, err := os.OpenFile(filePath+fileName, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n----> %s\n", err)
		return
	}
	defer outputf.Close()

	// 寫入Log
	// outputWriter := bufio.NewWriter(outputf)
}
