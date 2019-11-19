package middleware

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

/*
Request logging middleware configuration
*/
var r *rand.Rand

const (
	randTable string = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func RequestLogging(c *gin.Context) {
	// set trace id
	setTraceId(c)
	// log request information
	logReq(c.Request)
	// thought next
	c.Next()
}

// set trace id to track whole request lifecycle
func setTraceId(h *gin.Context) {
	h.Request.Header.Set("Trace-Id", genTraceId())
}

// logging request information
func logReq(req *http.Request) {
	//log.Println(req)
}

func genTraceId() string {

	bytes := []byte(randTable)
	var result []byte
	for i := 0; i < 32; i++ {
		if i > 0 && i%8 == 0 {
			result = append(result, '-')
		}
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
