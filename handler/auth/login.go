package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginOption 存放登入參數
type LoginOption struct {
	Username string
	Password string
}

// LoginResult 回傳登入結果
type LoginResult struct {
	Meta *LoginOption `json:"meta"`
	Data interface{}  `json:"data"`
}

// Login 登入功能
func Login(c *gin.Context) {
	l := &LoginOption{}
	l.Username = c.PostForm("username")
	l.Password = c.PostForm("password")

	lr := &LoginResult{}
	lr.Meta = l
	lr.Data = []string{}

	c.JSON(http.StatusOK, *lr)
}
