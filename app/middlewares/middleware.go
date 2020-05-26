package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
Middlewares Configuration
*/
type Middlewares []gin.HandlerFunc

var (
	MiddlewareChain Middlewares
)

func init() {
	fmt.Println("Middlewares initializing")
	// init middlewares
	MiddlewareChain = Middlewares{
		FilterInvalidRequest,
		RequestLogging,
	}
}
func InitMiddleware(engine *gin.Engine) {
	engine.Use(MiddlewareChain...)
}
