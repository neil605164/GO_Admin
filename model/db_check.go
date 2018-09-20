package model

import (
	"GO_Admin/global"
	"fmt"

	"github.com/jinzhu/gorm"
)

// checkUserTable 檢查users table 是否存在
func checkUserTable(tableName string, db *gorm.DB) (bool, error) {
	if !db.HasTable(tableName) {
		err := global.NewError{
			Title:   "table is not exist",
			Message: fmt.Sprintf("Users table is not exist, can not insert data"),
		}
		return false, err
	}

	return true, nil
}

// checkUserInfoTable 檢查users_infos table 是否存在
func checkUserInfoTable(tableName string, db *gorm.DB) (bool, error) {
	if !db.HasTable(tableName) {
		err := global.NewError{
			Title:   "table is not exist",
			Message: fmt.Sprintf("Users_Info table is not exist, can not insert data"),
		}
		return false, err
	}

	return true, nil
}

// CheckMemExist 檢查會員是否已經存在
func CheckMemExist(member string, db *gorm.DB) (bool, error) {
	var users []User

	// 不預期錯誤
	if err := db.Where("username = ?", member).Find(&users).Error; err != nil {
		err := global.NewError{
			Title:   "Unexpected error when check user exist",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return true, err
	}

	// 用戶已存在
	if len(users) > 0 {
		err := global.NewError{
			Title:   "Member is Exist",
			Message: fmt.Sprintf("%s member is exist", member),
		}
		return true, err
	}

	return false, nil
}
