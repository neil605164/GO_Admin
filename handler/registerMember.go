package member

import (
	_ "GO_Admin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerMember struct {
	Username string
	Password string
}

type registerMemberResult struct {
	Meta registerMember `json:"meta"`
	Data interface{}    `json:"data"`
}

func RegisterMember(c *gin.Context) {
	// dbConn.DBConnect()

	// 事先聲明defer,才可以抓到panic的值
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	log.Println("=======Register Start=======:")

	registerMember := &registerMember{}
	registerMember.Username = c.PostForm("account")
	registerMember.Password = c.PostForm("password")

	registerMemberResult := &registerMemberResult{}
	registerMemberResult.Meta = *registerMember

	c.JSON(http.StatusOK, registerMemberResult)
}
