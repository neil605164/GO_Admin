package model

import "time"

// User 定義 user model
// 備註:
// 1. AUTO_INCREMENT 必須用在 primary_key
type User struct {
	ID        int       `db:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	Username  string    `db:"username" gorm:"column:username;not null;unique;primary_key"`
	Password  string    `db:"password" gorm:"column:password;not null"`
	CreatedAt time.Time `db:"createdat" gorm:"column:createat"`
}
