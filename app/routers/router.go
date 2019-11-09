package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Routers initializing")
}
func InitRouters(engine *gin.Engine) {
	InitAccountRouter(engine)
	InitBookRouter(engine)
}
