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

	if CheckMemExist(rgMem.Username, db); err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	if err := tx.Create(&userInfo).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when register user",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	tx.Commit()
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
			Title:   "Unexpected error when get all users list",
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

	user := User{
		Username: edUserInfo.Username,
		Password: edUserInfo.Password,
	}

	userInfo := UserInfo{
		Username: edUserInfo.Username,
		Nickname: edUserInfo.Nickname,
		Email:    edUserInfo.Enail,
		Addr:     edUserInfo.Addr,
	}

	tx := db.Begin()

	if CheckMemExist(edUserInfo.Username, db); err != nil {
		return err
	}

	if err := tx.Model(&user).Where("username = ?", user.Username).Updates(&user).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when edit users table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	if err := tx.Model(&userInfo).Where("username = ?", userInfo.Username).Updates(&userInfo).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when edit user_infos table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	tx.Commit()
	return nil
}

// SQKFreezeUserAccount 停用用戶帳號
func SQKFreezeUserAccount(freezeMem *global.FreezeUserAccountOption) (err error) {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	user := User{
		Status: "1",
	}

	if CheckMemExist(freezeMem.Username, db); err != nil {
		return err
	}

	tx := db.Begin()

	if err = tx.Model(&user).Where("username = ?", freezeMem.Username).Update(&user).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when edit users table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}
	tx.Commit()

	return nil
}

// SQLDeleteUserAccount 刪除用戶帳號
func SQLDeleteUserAccount(deleteMem *global.DeleteUserAccountOption) (err error) {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	user := User{}
	userInfo := UserInfo{}

	// 可以移除，因為下方有檢查『影響數量』
	if CheckMemExist(deleteMem.Username, db); err != nil {
		return err
	}

	tx := db.Begin()
	execRes := tx.Model(&user).Where("users.username = ?", deleteMem.Username).Delete(&user)
	if execRes.Error != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when delete users table",
			Message: fmt.Sprintf("Error massage is: %s", execRes.Error),
		}
		return err
	}
	if execRes.RowsAffected == 0 {
		tx.Rollback()
		err = global.NewError{
			Title:   "data is not exist when delete users table",
			Message: fmt.Sprintf("Affected rows is: %d", execRes.RowsAffected),
		}
		return err
	}

	execRes = tx.Model(&userInfo).Where("username = ?", deleteMem.Username).Delete(&userInfo)
	if execRes.Error != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when delete user_infos table",
			Message: fmt.Sprintf("Error massage is: %s", execRes.Error),
		}
		return err
	}
	if execRes.RowsAffected == 0 {
		tx.Rollback()
		err = global.NewError{
			Title:   "data is not exist when delete user_infos table",
			Message: fmt.Sprintf("Affected rows is: %d", execRes.RowsAffected),
		}
		return err
	}

	tx.Commit()
	return nil
}

// SQLEnableUserAccount 啟用會員帳號
func SQLEnableUserAccount(enableMem *global.EnableUserAccountOption) (err error) {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	user := User{
		Status: "0",
	}

	if CheckMemExist(enableMem.Username, db); err != nil {
		return err
	}

	tx := db.Begin()

	if err = tx.Model(&user).Where("username = ?", enableMem.Username).Update(&user).Error; err != nil {
		tx.Rollback()
		err = global.NewError{
			Title:   "Unexpected error when edit users table",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}
	tx.Commit()

	return nil
}
