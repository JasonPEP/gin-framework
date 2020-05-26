package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
Filter invalid request for performance
*/
func FilterInvalidRequest(c *gin.Context) {
	url := c.Request.RequestURI
	if strings.HasPrefix(url, "/v1") {
		fmt.Println("Invalid Request")
		c.Abort()
	} else {
		c.Next()
	}
}
