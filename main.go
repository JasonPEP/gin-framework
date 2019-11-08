package main

import (
	"gin-framework/app/config"
	"gin-framework/app/routers"
	"net/http"
	"time"
)

func main() {
	// init database
	db := config.InitDB()
	defer db.Close()
	// init routers
	//routers.InitRouters().Run()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        routers.InitRouters(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()

}
