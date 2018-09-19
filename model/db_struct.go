/*******************************************
* 每次新增DB的Sturct，                      *
* 都需要調整dbcon.go的CheckTableIsExist，   *
* 進行Migrate驗證。                         *
*******************************************/

package model

import "time"

// User 定義 user model
// 備註:
// 存放會員帳密
// 1. AUTO_INCREMENT 必須用在 primary_key
// 2. status: 0:啟用的帳號 1:凍結的帳號 2:刪除的帳號
type User struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Username  string    `gorm:"column:username;not null;unique;primary_key"`
	Password  string    `gorm:"column:password;not null"`
	Status    int       `gorm:"column:status;default:0"`
	CreatedAt time.Time `gorm:"column:createat"`
}

// UserInfo 定義 userinfo model
// 備註:
// 存放會員詳細資料
type UserInfo struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Username  string    `gorm:"column:username;not null;unique;primary_key"`
	Nickname  string    `gorm:"column:nickname;"`
	Email     string    `gorm:"column:email;"`
	Addr      string    `gorm:"column:addr;"`
	CreatedAt time.Time `gorm:"column:createat"`
}
