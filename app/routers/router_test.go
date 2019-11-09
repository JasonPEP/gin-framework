package router_test

import (
	"gin-framework/app"
	router "gin-framework/app/routers"
	"testing"
)

func TestInitAccountRouter(t *testing.T) {
	router.InitRouters(app.Engine)
}
