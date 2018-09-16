package model

// User 定義 user table
type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	// CreatedAt time.Time `db:"createdat"`
}
