package main

import (
	"fmt"
	"go-blog-demo/pkg/setting"
	"go-blog-demo/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:		   fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:		router,
		ReadTimeout:	setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
