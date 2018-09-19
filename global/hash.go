package global

import (
	"crypto/md5"
	"fmt"
)

// Md5Encryption md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}
