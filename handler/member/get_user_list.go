package member

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserList 取得使用者清單
func GetUserList(c *gin.Context) {
	log.Println("=======Get User List Start=======:")
	// compose param start
	getUserListResult := global.GetUserListResult{}

	// execute db start
	dbResult, err := model.SQLGetUserList()
	if err != nil {
		getUserListResult.Data = err
	}
	// execute db end

	// 若不定義成空陣列，會顯示null字串
	getUserListResult.Meta = []string{}
	getUserListResult.Data = dbResult
	// compose param end
	log.Println("=======Get User List End=======:")

	c.JSON(http.StatusOK, getUserListResult)
}
