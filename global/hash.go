package global

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

// Md5Encryption md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}

// Md5EncryptionWithTime md5 加密（加上時間）
func Md5EncryptionWithTime(str string) string {
	crutime := time.Now().Unix()
	data := str + strconv.FormatInt(crutime, 10)
	key := []byte(data)

	token := md5.Sum(key)
	md5Str := fmt.Sprintf("%x", token)

	return md5Str
}
