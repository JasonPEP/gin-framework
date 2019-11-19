package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

/*
Middlewares Configuration
*/
type Middlewares []gin.HandlerFunc

var (
	chain Middlewares
)

func init() {
	log.Println("Middlewares initializing")
	// init middlewares
	chain = Middlewares{
		FilterInvalidRequest,
		RequestLogging,
		Auth,
	}
}
func InitMiddleware(engine *gin.Engine) {
	engine.Use(chain...)
}
