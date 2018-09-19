package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
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
	registerMemberOption.Nickname = c.PostForm("nickname")
	registerMemberOption.Enail = c.PostForm("email")
	registerMemberOption.Addr = c.PostForm("addr")
	// get param ecd

	// encryption password start
	registerMemberOption.Password = global.Md5Encryption(registerMemberOption.Password)
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
