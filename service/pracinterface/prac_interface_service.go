package pracInterface

import (
	"fmt"
	"strconv"
)

// Database 宣告interface
type Database interface {
	Read(name string) string
	Write(value int) string
}

// Paper 模擬紙張
type Paper struct {
}

// New 呼叫interface
func New(db Database) {
	db.Read("Neil")
	db.Write(123)
}

func (p Paper) Read(name string) string {
	newName := name
	fmt.Println(newName)
	return newName
}

func (p Paper) Write(value int) string {
	newValue := strconv.Itoa(value)
	fmt.Println(newValue)
	return newValue
}
