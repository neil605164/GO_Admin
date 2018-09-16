package model

type db struct {
	Title   string
	Message string
}

type user struct {
	Username string `db:"username"`
	Password string `db:"password"`
	// CreatedAt time.Time `db:"createdat"`
}
