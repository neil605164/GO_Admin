package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditUserInfo(c *gin.Context) {
	c.String(http.StatusOK, "edit user info")
}
