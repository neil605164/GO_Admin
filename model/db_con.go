package model

import (
	"GO_Admin/global"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// dbConnect 建立 DB 連線
func dbConnect() (db *gorm.DB) {
	USER := global.Config.Database.User
	PASSWORD := global.Config.Database.Password
	HOST := global.Config.Database.Host
	DATABASE := global.Config.Database.Database

	// 組合連線資訊
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, DATABASE)

	// 建立連線
	db, err := gorm.Open("mysql", connectionString)
	checkError(err)

	return db
}

// CheckTableIsExist 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func CheckTableIsExist() {
	db := dbConnect()

	defer db.Close()

	if !db.HasTable("users") {
		db.AutoMigrate(&User{})
	}
}

// SQLRegisterMem 註冊會員
func SQLRegisterMem(rgMem *global.RegisterMemberOption) (err error) {
	user := User{
		Username: rgMem.Username,
		Password: rgMem.Password,
		// CreatedAt: time.Now(),
	}

	db := dbConnect()
	defer db.Close()

	// 檢查DB是否存在，若存在才可以新增，否則回傳錯誤
	if !db.HasTable("users") {
		err = global.NewError{
			Title:   "table is not exist",
			Message: fmt.Sprintf("Users table is not exist, can not insert data"),
		}
		return err
	}

	// 檢查(主鍵)資料是否已經存在
	// isExist := db.NewRecord(user)
	// fmt.Printf("=========%v=========", isExist)
	// if isExist {
	// 	err = global.NewError{
	// 		Title:   "Data is Exist !!!!!",
	// 		Message: fmt.Sprintf("%s member is exist", user.Username),
	// 	}

	// 	return err
	// }

	err = db.Create(&user).Error
	if err != nil {
		err = global.NewError{
			Title:   "Member is Exist",
			Message: fmt.Sprintf("%s member is exist", user.Username),
		}
		return err
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
