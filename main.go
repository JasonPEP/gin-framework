package main

import (
	"gin-framework/app"
)

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
	//s := &http.Server{
	//	Addr:           ":8888",
	//	Handler:        routers.InitRouters(),
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//_ = s.ListenAndServe()
}
