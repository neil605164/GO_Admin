package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableUserAccount(c *gin.Context) {
	// get param start
	enableUserAccountOption := &global.EnableUserAccountOption{}
	enableUserAccountOption.Username = c.PostForm("account")
	// get param end

	// compose param start
	enableUserAccountResult := &global.EnableUserAccountResult{}
	enableUserAccountResult.Meta = *enableUserAccountOption

	// execute db start
	if err := model.SQLEnableUserAccount(enableUserAccountOption); err != nil {
		enableUserAccountResult.Data = err
		c.JSON(http.StatusOK, enableUserAccountResult)
		return
	}
	// execute db end

	enableUserAccountResult.Data = "Enable user account access"
	// compose param end

	c.JSON(http.StatusOK, enableUserAccountResult)
}
