package file

import (
	"GO_Admin/global"
	"GO_Admin/service/file"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile 上傳檔案
func UploadFile(c *gin.Context) {
	log.Println("=======Upload File Start=======:")
	uploadFileRes := &global.UploadFileResult{}

	// get param
	params, err := fileservice.ComposeParams(c)
	if err != nil {
		uploadFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadFileRes)
		return
	}

	uploadFileRes.Meta = params.File

	// check path exist
	if err := fileservice.CheckAndMakeDir(params.FilePath); err != nil {
		uploadFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadFileRes)
		return
	}

	// create file ＆ insert to files table
	if err = fileservice.CreateFile(c, params); err != nil {
		uploadFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadFileRes)
		return
	}

	uploadFileRes.Data = "Create file success"

	// return API
	c.JSON(http.StatusOK, uploadFileRes)
}
