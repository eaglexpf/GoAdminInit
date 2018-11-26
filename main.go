package main

import (
	"fmt"
	"net/http"

	"github.com/eaglexpf/GoAdminInit/pkg/setting"

	"github.com/eaglexpf/GoAdminInit/routers"
)

func main() {
	router := routers.InitRouter()
	fmt.Println("main:", setting.HttpPort)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
