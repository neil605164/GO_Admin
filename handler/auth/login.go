package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginOption 存放登入參數
type LoginOption struct {
	username string
	password string
}

// LoginResult 回傳登入結果
type LoginResult struct {
	Meta LoginOption `json:"meta"`
	Data interface{} `json:"data"`
}

// Login 登入功能
func Login(c *gin.Context) {
	l := &LoginOption{}
	l.username = c.PostForm("username")
	l.password = c.PostForm("password")

	lr := &LoginResult{}

	fmt.Println(*l)
	lr.Meta = *l
	lr.Data = "123"

	fmt.Println(lr.Meta)
	c.JSON(http.StatusOK, *lr)
}
