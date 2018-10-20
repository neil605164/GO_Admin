package route

import (
	"GO_Admin/handler/file"
	"GO_Admin/handler/member"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由控制
func SetupRouter(r *gin.Engine) {

	// 登入驗證
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	// 註冊會員
	authorized.POST("/register_member", member.RegisterMember)
	// 取得會員清單
	authorized.GET("/get_user_list", member.GetUserList)
	// 編輯會員資料
	authorized.PUT("/edit_user_info", member.EditUserInfo)
	// 停用會員帳號
	authorized.PUT("/freeze_user_account", member.FreezeUserAccount)
	// 刪除會員帳號
	authorized.DELETE("/delete_user_account/:account", member.DeleteUserAccount)
	// 啟用會員帳號
	authorized.PUT("/enable_user_account", member.EnableUserAccount)

	// 上傳檔案(單一)
	authorized.POST("/upload_file", file.UploadFile)
	// 上傳檔案(多張)
	authorized.POST("/upload_multifile", file.UploadMultiFile)
}
