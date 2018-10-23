package pracInterface

import (
	"GO_Admin/service/pracinterface"

	"github.com/gin-gonic/gin"
)

// InterfacePrac 練習interface
func InterfacePrac(c *gin.Context) {
	p := pracInterface.Paper{}
	pracInterface.New(&p)
}
