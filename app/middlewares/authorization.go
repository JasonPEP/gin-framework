package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	AuthKey string = "Authorization"
)

func Auth(c *gin.Context) {
	a := c.GetHeader(AuthKey)
	if a == "" {
		a := errors.New("Unauthorizative")
		_ = c.AbortWithError(http.StatusUnauthorized, a)
	}
	c.Next()
}
