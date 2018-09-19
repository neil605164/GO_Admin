package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FreezeUserAccount(c *gin.Context) {
	c.String(http.StatusOK, "freeze user account")
}
