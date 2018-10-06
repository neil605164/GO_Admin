package fileservice

import (
	"GO_Admin/global"
	"GO_Admin/model"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetFileSize 取檔案大小
func GetFileSize(file *multipart.FileHeader) int64 {
	// 取得尺寸後，轉換成km或mb
	return file.Size
}

// GetFileExt 取檔案副檔名(移除「.」符號)
func GetFileExt(file *multipart.FileHeader) string {
	fileName := file.Filename
	fileExt := strings.Replace(filepath.Ext(fileName), ".", "", -1)
	return fileExt
}

// CheckPermission 檢查檔案權限
func CheckPermission(src string) (bool, error) {
	return true, nil
}

// CheckAndMakeDir 檢查資料夾是否存在，若不存在新增資料夾
func CheckAndMakeDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0755); err != nil {
			err = global.NewError{
				Title:   "Unexpected error when mkdir new dir",
				Message: fmt.Sprintf("Error massage is: %s", err),
			}
			return err
		}
	}
	return nil
}

// ComposeParams 組合上傳檔案參數
func ComposeParams(c *gin.Context) (*global.UploadFileOption, error) {
	uploadOption := &global.UploadFileOption{}
	file, err := c.FormFile("file")
	if err != nil {
		err = global.NewError{
			Title:   "Unexpected error when upload file",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return uploadOption, err
	}

	// 定義檔案路徑
	filePath := "/home/hsieh/go/src/GO_Admin/File"

	uploadOption.File = file
	uploadOption.FileName = global.Md5EncryptionWithTime(file.Filename)
	uploadOption.FileSize = GetFileSize(file)
	uploadOption.FileExt = GetFileExt(file)
	uploadOption.FilePath = filePath

	return uploadOption, nil
}

// CreateFile 建立新檔案
func CreateFile(c *gin.Context, params *global.UploadFileOption) error {
	// 組合路徑
	path := params.FilePath + "/" + params.FileName

	// 建立檔案
	if err := c.SaveUploadedFile(params.File, path); err != nil {
		err = global.NewError{
			Title:   "Unexpected error when create new file",
			Message: fmt.Sprintf("Error massage is: %s", err),
		}
		return err
	}

	// 存入資料庫
	if err := model.SQLUploadFile(params); err != nil {
		return err
	}

	return nil
}
