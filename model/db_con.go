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
		db.AutoMigrate(&User{})
	}

	if !db.HasTable("user_infos") {
		db.AutoMigrate(&UserInfo{})
	}

	return nil
}

// SQLRegisterMem 註冊會員
func SQLRegisterMem(rgMem *global.RegisterMemberOption) (err error) {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	errorMsg := make(chan error)
	user := User{
		Username: rgMem.Username,
		Password: rgMem.Password,
	}

	userInfo := UserInfo{
		Username: rgMem.Username,
		Nickname: rgMem.Nickname,
		Email:    rgMem.Enail,
		Addr:     rgMem.Addr,
	}

	if checkUserTable("users", db); err != nil {
		return err
	}

	if checkUserInfoTable("user_infos", db); err != nil {
		return err
	}

	if CheckMemExist(rgMem.Username, db); err != nil {
		return err
	}

	go func(user *User, db *gorm.DB) {
		err = createUserData(user, db)
		errorMsg <- err
	}(&user, db)

	go func(userInfo *UserInfo, db *gorm.DB) {
		err = createUserInfoData(userInfo, db)
		errorMsg <- err
	}(&userInfo, db)

	err = <-errorMsg
	return err
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

// SQLEditUserInfo 編輯會員資料
func SQLEditUserInfo(edUserInfo *global.EditUserInfoOption) (err error) {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	errorMsg := make(chan error)

	user := User{
		Username:  edUserInfo.Username,
		Password:  edUserInfo.Password,
		CreatedAt: time.Now(),
	}

	userInfo := UserInfo{
		Username:  edUserInfo.Username,
		Nickname:  edUserInfo.Nickname,
		Email:     edUserInfo.Enail,
		Addr:      edUserInfo.Addr,
		CreatedAt: time.Now(),
	}

	if checkUserTable("users", db); err != nil {
		return err
	}

	if checkUserInfoTable("user_infos", db); err != nil {
		return err
	}

	if CheckMemExist(edUserInfo.Username, db); err != nil {
		return err
	}

	go func(user *User, db *gorm.DB) {
		err = updateUserData(user, db)
		errorMsg <- err
	}(&user, db)

	go func(userInfo *UserInfo, db *gorm.DB) {
		err = updateUserInfoData(userInfo, db)
		errorMsg <- err
	}(&userInfo, db)

	err = <-errorMsg

	return nil
}

// createUserData 新增資料至users table
func createUserData(user *User, db *gorm.DB) error {
	if err := db.Create(&user).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	return nil
}

// createUserInfoData 新增資料至user_infos table
func createUserInfoData(userInfo *UserInfo, db *gorm.DB) error {
	if err := db.Create(&userInfo).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}
	return nil
}

// updateUserData 更新資料至users table
func updateUserData(user *User, db *gorm.DB) error {
	if err := db.Model(&user).Where("username = ?", user.Username).Updates(&user).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when edit users table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}
	return nil
}

// updateUserInfoData 新增資料至user_infos table
func updateUserInfoData(edUserInfo *UserInfo, db *gorm.DB) error {
	if err := db.Model(&edUserInfo).Where("username = ?", edUserInfo.Username).Updates(&edUserInfo).Error; err != nil {
		err = global.NewError{
			Title:   "Unexpected error when edit user_infos table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}
	return nil
}
