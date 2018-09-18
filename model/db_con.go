package model

import (
	"GO_Admin/global"
	"fmt"

	"github.com/jinzhu/gorm"
)

// dbConnect 建立 DB 連線
func dbConnect() (db *gorm.DB, err error) {
	USER := global.Config.Database.User
	PASSWORD := global.Config.Database.Password
	HOST := global.Config.Database.Host
	DATABASE := global.Config.Database.Database

	// 組合連線資訊
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, DATABASE)

	// 建立連線
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		err = global.NewError{
			Title:   "DB connect Fail",
			Message: fmt.Sprintf("Error message is: %s", err),
		}
		return nil, err
	}

	return db, nil
}

// CheckTableIsExist 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func CheckTableIsExist() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	if !db.HasTable("users") {
		db.AutoMigrate(&User{})
	}

	if !db.HasTable("user_infos") {
		db.AutoMigrate(&User_Info{})
	}

	return nil
}

// SQLRegisterMem 註冊會員
func SQLRegisterMem(rgMem *global.RegisterMemberOption) (err error) {
	user := User{
		Username: rgMem.Username,
		Password: rgMem.Password,
	}

	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	// 檢查DB是否存在，若存在才可以新增，否則回傳錯誤
	if !db.HasTable("users") {
		err = global.NewError{
			Title:   "table is not exist",
			Message: fmt.Sprintf("Users table is not exist, can not insert data"),
		}
		return err
	}

	// 檢查會員是否已存在
	memExist, err := CheckMemExist(rgMem.Username, db)
	if err != nil {
		err = global.NewError{
			Title:   "Unexpected error when check user exist",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	if memExist {
		err = global.NewError{
			Title:   "Member is Exist",
			Message: fmt.Sprintf("%s member is exist", user.Username),
		}
		return err
	}

	if err = db.Create(&user).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	return nil
}

// SQLGetUserList 取得用戶清單
func SQLGetUserList() (userList *[]User, err error) {
	var users []User

	db, err := dbConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Find(&users).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when get all user list",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return nil, err
	}
	fmt.Println(&users)
	return &users, nil
}

// CheckMemExist 檢查會員是否已經存在
func CheckMemExist(member string, db *gorm.DB) (bool, error) {
	var users []User

	// 不預期錯誤
	if err := db.Where("username = ?", member).Find(&users).Error; err != nil {
		return true, err
	}

	// 用戶已存在
	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}
