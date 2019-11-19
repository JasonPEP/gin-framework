package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ResponseLogging(c *gin.Context) {
	c.Next()
	// log response information
	logRw(c.Writer)

}

// logging response information
func logRw(rw gin.ResponseWriter) {
	log.Println(rw.Status())
}
