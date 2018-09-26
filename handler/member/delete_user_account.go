package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUserAccount 刪除使用者帳號
func DeleteUserAccount(c *gin.Context) {
	// get parameter start
	deleteUserAccountOption := &global.DeleteUserAccountOption{}
	deleteUserAccountOption.Username = c.Params.ByName("account")
	// get parameter end

	// compose result start
	deleteUserAccountResult := &global.DelteUserAccountResult{}
	deleteUserAccountResult.Meta = *deleteUserAccountOption

	// execute db start
	if err := model.SQLDeleteUserAccount(deleteUserAccountOption); err != nil {
		deleteUserAccountResult.Data = err
		c.JSON(http.StatusOK, deleteUserAccountResult)
		return
	}
	// execute db end

	deleteUserAccountResult.Data = "刪除用戶帳號成功"
	// compose result end
	c.JSON(http.StatusOK, deleteUserAccountResult)
}
