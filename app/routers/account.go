package routers

import (
	"gin-framework/app/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountRouter(engine *gin.Engine) {
	group := engine.Group("account")
	{
		group.GET("account", controller.GetAccount)
	}

}
