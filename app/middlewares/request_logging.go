package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
Request logging middleware configuration
*/

func RequestLogging(c *gin.Context) {
	fmt.Println("Request begin")
	c.Next()
	fmt.Println("Request over")
}
