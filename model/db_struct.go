package model

type DB struct {
	Title   string
	Message string
}

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	// CreatedAt time.Time `db:"createdat"`
}
