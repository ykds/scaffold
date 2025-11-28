package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"scaffold/config"
	"scaffold/handler"
	"scaffold/middleware"
	"scaffold/pkg/logger"
	"scaffold/pkg/mongodb"
	"scaffold/pkg/redis"
	"scaffold/pkg/tdengine"
	"scaffold/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("c", "./config.yaml", "set the config.yaml path")

func main() {
	flag.Parse()

	// 初始化配置
	cfg := config.InitConfig(*configFile)
	logger.InitLogger(cfg.Logger)
	mgo := mongodb.NewMongo(cfg.Mongo)
	defer mgo.Close()
	rdb := redis.NewRedis(cfg.Redis)
	defer rdb.Close()
	taos := tdengine.NewTaos(cfg.Taos)
	defer taos.Close()

	// 初始化 http server
	engine := gin.New()
	engine.Use(middleware.GinLogger(), gin.RecoveryWithWriter(logger.GetOutput()))
	gin.SetMode(gin.ReleaseMode)
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	}
	services := service.NewServices(mgo, rdb, taos)
	handler.InitHandlers(services)
	handler.RegisterRouter(engine)

	server := http.Server{
		Addr:         cfg.Server.Port,
		Handler:      engine,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	// 程序结束处理
	done := make(chan error)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			done <- err
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		logger.Infof("program exit with signal")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	case err := <-done:
		logger.Errorf("program exit. err: %v", err)
	}
}
