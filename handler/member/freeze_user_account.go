package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FreezeUserAccount 停用使用者帳號
func FreezeUserAccount(c *gin.Context) {
	log.Println("=======Freeze User Acount Start=======:")
	// get param start
	freezeUserAccountOption := &global.FreezeUserAccountOption{}
	freezeUserAccountOption.Username = c.PostForm("account")
	// get param end

	// compose param start
	freezeUserAccountResult := &global.FreezeUserAccountResult{}
	freezeUserAccountResult.Meta = *freezeUserAccountOption

	// execute db start
	err := model.SQKFreezeUserAccount(freezeUserAccountOption)
	if err != nil {
		freezeUserAccountResult.Data = err
		c.JSON(http.StatusOK, freezeUserAccountResult)
		return
	}
	// execute db end

	freezeUserAccountResult.Data = "Freeze user account access"
	// compose param end

	c.JSON(http.StatusOK, freezeUserAccountResult)
}
