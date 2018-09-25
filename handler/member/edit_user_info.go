package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EditUserInfo 編輯使用者資訊
func EditUserInfo(c *gin.Context) {
	log.Println("=======Edit User Info Start=======:")
	// get param start
	editUserInfoOption := &global.EditUserInfoOption{}
	editUserInfoOption.Username = c.PostForm("account")
	editUserInfoOption.Password = c.PostForm("password")
	editUserInfoOption.Nickname = c.PostForm("nickname")
	editUserInfoOption.Enail = c.PostForm("email")
	editUserInfoOption.Addr = c.PostForm("addr")
	// get param end

	// encryption password start
	editUserInfoOption.Password = global.Md5Encryption(editUserInfoOption.Password)
	// encryption password end

	// compose param start
	editUserInfoResult := &global.EditUserInfoResult{}
	editUserInfoResult.Meta = *editUserInfoOption

	// execute db start
	err := model.SQLEditUserInfo(editUserInfoOption)
	if err != nil {
		editUserInfoResult.Data = err
		c.JSON(http.StatusOK, editUserInfoResult)
		return
	}
	// execute db end

	editUserInfoResult.Data = "Access Edit User Info"
	// compose param end

	c.JSON(http.StatusOK, *editUserInfoResult)
}
