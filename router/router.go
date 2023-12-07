package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenzh1990/shell-probe/middleware"
	"github.com/shenzh1990/shell-probe/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func SetupRouter(port int) {
	r := initRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Info(fmt.Sprintf("Listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Info("Server exiting")

}
func initRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors()) //配置跨域
	server := r.Group("/api")
	server.Use(middleware.AutuWithOutJwtMiddelware())
	{
		server.GET("/cpu", service.GetCPUInfo)
		server.GET("/mem", service.GetMemInfo)
		server.GET("/disk", service.GetDiskInfo)
		server.GET("/net", service.GetNetInfo)
		server.GET("/host", service.GetHostInfo)
		server.GET("/load", service.GetLoadInfo)
	}
	//配置密码
	ws := r.Group("/ws")
	{
		ws.GET("/host", service.GetWsHostInfo)
	}
	return r
}
