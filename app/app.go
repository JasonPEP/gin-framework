package app

import (
	"fmt"
	middleware "gin-framework/app/middlewares"
	model "gin-framework/app/models"
	router "gin-framework/app/routers"
	"github.com/gin-gonic/gin"
)

var (
	Engine *gin.Engine
)

func init() {
	fmt.Println("Application initializing")
	Engine = gin.New()
}
func Run() (err error) {
	model.InitDatabase()
	// init middlewares
	middleware.InitMiddleware(Engine)
	// init routers
	router.InitRouters(Engine)
	return Engine.Run(":8080")
}
