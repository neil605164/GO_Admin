/*******************************************
* 每次新增DB的Sturct，                      *
* 都需要調整dbcon.go的CheckTableIsExist，   *
* 進行Migrate驗證。                         *
*******************************************/

package model

import "github.com/jinzhu/gorm"

// User 定義 user model
// 備註:
// 存放會員帳密
// 1. AUTO_INCREMENT 必須用在 primary_key
// 2. status: 0:啟用的帳號 1:凍結的帳號
type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null;unique;primary_key"`
	Password string `gorm:"column:password;not null"`
	Status   int    `gorm:"column:status;default:0"`
}

// UserInfo 定義 userinfo model
// 備註:
// 存放會員詳細資料
type UserInfo struct {
	gorm.Model
	Username string `gorm:"column:username;not null;unique;primary_key"`
	Nickname string `gorm:"column:nickname;"`
	Email    string `gorm:"column:email;"`
	Addr     string `gorm:"column:addr;"`
}

// File 定義files table
// 存放上傳檔案的詳細資料
type File struct {
	gorm.Model
	FileName string `gorm:"column:file_name;not null"`
	FilePath string `gorm:"column:file_path;not null"`
	FileSize int64  `gorm:"column:file_size"`
	FileExt  string `gorm:"column:file_ext"`
}
