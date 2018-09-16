package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMember(c *gin.Context) {
	// 事先聲明defer,才可以抓到panic的值
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	log.Println("=======Register Start=======:")
	// get param start
	registerMemberOption := &global.RegisterMemberOption{}
	registerMemberOption.Username = c.PostForm("account")
	registerMemberOption.Password = c.PostForm("password")
	// get param ecd

	// compose param start
	registerMemberResult := &global.RegisterMemberResult{}
	registerMemberResult.Meta = *registerMemberOption

	// execute db start
	err := model.SQL_RegisterMem(registerMemberOption)
	if err != nil {
		registerMemberResult.Data = err
		fmt.Printf("=========%v=========", err)
		c.JSON(http.StatusOK, registerMemberResult)
		return
	}
	// execute db end

	registerMemberResult.Data = "Access Register Member"
	// compose param end

	c.JSON(http.StatusOK, *registerMemberResult)
}
