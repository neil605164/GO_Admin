package model

import "time"

// User 定義 user model
type User struct {
	ID        int       `db:"id" gorm:"primary_key"`
	Username  string    `db:"username" gorm:"column:username;not null;unique"`
	Password  string    `db:"password" gorm:"column:password;not null"`
	CreatedAt time.Time `db:"createdat" gorm:"column:createat"`
}
