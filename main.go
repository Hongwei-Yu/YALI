package main

import (
	"YALI/initialize"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var GinServer *gin.Engine

func initservice() {
	GinServer = initialize.Routers()
	kpRunnerService := &http.Server{
		Addr:           "localhost:8080",
		Handler:        GinServer,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := kpRunnerService.ListenAndServe(); err != nil {
			log.Error(fmt.Sprintf("机器ip:%s, kpRunnerService:", "127.0.0.1"), err)
			return
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(fmt.Sprintf("机器ip:%s, 注销成功", "127.0.0.1"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := kpRunnerService.Shutdown(ctx); err != nil {
		log.Info(fmt.Sprintf("机器ip:%s, 注销成功", "127.0.0.1"))
	}
}

func main() {
	initservice()
}
