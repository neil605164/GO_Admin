package global

import "fmt"

// const (
// 	REGISTER_MEMBER = "INSERT INTO "
// )

type NewError struct {
	Title   string
	Message string
}

type DevConfig struct {
	Database Dbconnect `yaml:"database"`
}

type Dbconnect struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RegisterMemberOption struct {
	Username string
	Password string
}
type RegisterMemberResult struct {
	Meta RegisterMemberOption `json:"meta"`
	Data interface{}          `json:"data"`
}

func (e NewError) Error() string {
	return fmt.Sprintf("%v: %v", e.Title, e.Message)
}
