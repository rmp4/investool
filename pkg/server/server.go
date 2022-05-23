package server

import (
	"fmt"
	"investool/pkg/config"
	"investool/pkg/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() error {
	gin.SetMode(config.Get().Server.RunMode)

	routersInit := routes.InitRouter()
	readTimeout := config.Get().Server.ReadTimeout * time.Second
	writeTimeout := config.Get().Server.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", config.Get().Server.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	return server.ListenAndServe()
}
