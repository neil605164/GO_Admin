package model

import (
	"GO_Admin/global"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// dbConnect 建立 DB 連線
func dbConnect() (db *gorm.DB, err error) {
	USER := global.Config.Database.User
	PASSWORD := global.Config.Database.Password
	HOST := global.Config.Database.Host
	DATABASE := global.Config.Database.Database

	// 組合連線資訊
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, HOST, DATABASE)

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
		db.AutoMigrate(&user{})
	}

	if !db.HasTable("user_infos") {
		db.AutoMigrate(&userInfo{})
	}

	return nil
}

// SQLRegisterMem 註冊會員
func SQLRegisterMem(rgMem *global.RegisterMemberOption) (err error) {
	user := user{
		Username: rgMem.Username,
		Password: rgMem.Password,
	}

	userInfo := userInfo{
		Username: rgMem.Username,
		Nickname: rgMem.Nickname,
		Email:    rgMem.Enail,
		Addr:     rgMem.Addr,
	}

	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	checkUserBool, err := checkUserTable("users", db)
	if !checkUserBool && err != nil {
		return err
	}

	checkUserInfoBool, err := checkUserInfoTable("user_infos", db)
	if !checkUserInfoBool && err != nil {
		return err
	}

	memExistBool, err := CheckMemExist(rgMem.Username, db)
	if memExistBool && err != nil {
		return err
	}

	// if err = db.Create(&user).Error; err != nil {
	// 	err = global.NewError{
	// 		Title:   "Unexpected error when register user",
	// 		Message: fmt.Sprintf("Error massage is: %s", err),
	// 	}
	// 	return err
	// }
	go func(user *user, db *gorm.DB) {
		createMemberData(&user, db)
	}(&user, db)

	if err = db.Create(&userInfo).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	return nil
}

// SQLGetUserList 取得用戶清單
func SQLGetUserList() (userList *[]user, err error) {
	var users []user

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

// SQLEditUserInfo 編輯會員資料
func SQLEditUserInfo(edUserInfo *global.EditUserInfoOption) (err error) {
	user := user{
		Username:  edUserInfo.Username,
		Password:  edUserInfo.Password,
		CreatedAt: time.Now(),
	}

	userInfo := userInfo{
		Username:  edUserInfo.Username,
		Nickname:  edUserInfo.Nickname,
		Email:     edUserInfo.Enail,
		Addr:      edUserInfo.Addr,
		CreatedAt: time.Now(),
	}

	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	checkUserBool, err := checkUserTable("users", db)
	if !checkUserBool && err != nil {
		return err
	}

	checkUserInfoBool, err := checkUserInfoTable("user_infos", db)
	if !checkUserInfoBool && err != nil {
		return err
	}

	memExistBool, err := CheckMemExist(edUserInfo.Username, db)
	if memExistBool && err != nil {
		return err
	}

	if err = db.Model(&user).Where("username = ?", user.Username).Updates(&user).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when edit users table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	if err = db.Model(&userInfo).Where("username = ?", edUserInfo.Username).Updates(&userInfo).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when edit user_infos table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	return nil
}

func createMemberData(user *user, db *gorm.DB) error {
	if err := db.Create(&user).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	return nil
}
