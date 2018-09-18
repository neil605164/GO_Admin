package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterMember 註冊會員 start
func RegisterMember(c *gin.Context) {
	log.Println("=======Register Start=======:")
	// get param start
	registerMemberOption := &global.RegisterMemberOption{}
	registerMemberOption.Username = c.PostForm("account")
	registerMemberOption.Password = c.PostForm("password")
	// get param ecd

	// encryption password start
	registerMemberOption.Password = md5Encryption(registerMemberOption.Password)
	// encryption password end

	// compose param start
	registerMemberResult := &global.RegisterMemberResult{}
	registerMemberResult.Meta = *registerMemberOption

	// execute db start
	err := model.SQLRegisterMem(registerMemberOption)
	if err != nil {
		registerMemberResult.Data = err
		c.JSON(http.StatusOK, registerMemberResult)
		return
	}
	// execute db end

	registerMemberResult.Data = "Access Register Member"
	// compose param end

	c.JSON(http.StatusOK, *registerMemberResult)
}

func md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}
