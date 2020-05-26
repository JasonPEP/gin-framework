package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsCouting = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_http_api_stat",
		Help: "The total number of processed events",
	})
	opsConsuming = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_http_api_time",
		Help: "The total number of processed events",
	})
)

func RequestTracing(c *gin.Context) {
	couting(c)
	begin := time.Now()
	c.Next()
	end := time.Now().Sub(begin)
	consuming(c, end)
	log.Fatal(end.Milliseconds())
}
func consuming(c *gin.Context, t time.Duration) {
	go func() {
		opsConsuming.Inc()

	}()

}
func couting(c *gin.Context) {
	go func() {
		opsCouting.Inc()

	}()
}
