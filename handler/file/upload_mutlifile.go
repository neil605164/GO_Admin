package file

import (
	"GO_Admin/global"
	"GO_Admin/service/file"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadMultiFile 上傳檔案
func UploadMultiFile(c *gin.Context) {
	log.Println("=======Upload Mutli File Start=======:")
	uploadMultiFileRes := &global.UploadMultiFileResult{}

	// get param
	params, err := fileservice.ComposeMultiParams(c)
	if err != nil {
		uploadMultiFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadMultiFileRes)
		return
	}

	uploadMultiFileRes.Meta = params.File

	// check path exist
	if err := fileservice.CheckAndMakeDir(params.FilePath); err != nil {
		uploadMultiFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadMultiFileRes)
		return
	}

	// create file ＆ insert to files table
	if err = fileservice.CreateMultiFile(c, params); err != nil {
		uploadMultiFileRes.Data = err
		c.JSON(http.StatusBadRequest, uploadMultiFileRes)
		return
	}

	// uploadMultiFileRes.Data = "Create file success"

	// return API
	// c.JSON(http.StatusOK, uploadMultiFileRes)
}
