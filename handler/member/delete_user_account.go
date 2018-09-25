package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUserAccount 刪除使用者帳號
func DeleteUserAccount(c *gin.Context) {
	c.String(http.StatusOK, "delete user account")
}
