package app

import (
	"fmt"
	middleware "gin-framework/app/middlewares"
	router "gin-framework/app/routers"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func init() {
	fmt.Println("Application initializing")
	engine = gin.New()
}
func Run() (err error) {
	// init middlewares
	middleware.InitMiddleware(engine)
	// init routers
	router.InitRouters(engine)
	return engine.Run(":8080")
}
