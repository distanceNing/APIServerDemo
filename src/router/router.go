package router



import (
	"github.com/gin-gonic/gin"
	"../service"
)

func InitRouter() ( *gin.Engine) {
	r := gin.Default()

	r.GET("/",service.Handler)

	return r
}
