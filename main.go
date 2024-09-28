package main

import (
	"YALI/initialize"
	"YALI/log"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var GinServer *gin.Engine

func initservice() {
	log.InitLogger()
	GinServer = initialize.Routers()
	kpRunnerService := &http.Server{
		Addr:           "localhost:8080",
		Handler:        GinServer,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := kpRunnerService.ListenAndServe(); err != nil {
			log.Logger.Error(fmt.Sprintf("机器ip:%s, kpRunnerService:", "127.0.0.1"), err)
			return
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info(fmt.Sprintf("机器ip:%s, 注销成功", "127.0.0.1"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := kpRunnerService.Shutdown(ctx); err != nil {
		log.Logger.Info(fmt.Sprintf("机器ip:%s, 注销成功", "127.0.0.1"))
	}

}

func main() {
	initservice()
}
