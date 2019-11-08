package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	InitAccountRouter(Router)

	return Router
}
