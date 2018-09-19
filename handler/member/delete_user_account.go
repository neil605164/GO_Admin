package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserAccount(c *gin.Context) {
	c.String(http.StatusOK, "delete user account")
}
