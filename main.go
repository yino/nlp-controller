package main

import (
	"fmt"
	"os"

	"github.com/yino/nlp-controller/interfaces/middleware"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/config/router"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// load config
	config.GetConf()
	os.Setenv("GIN_MODE", "debug")
	gin.SetMode(gin.DebugMode)
	app := gin.New()
	pprof.Register(app)
	// 注册路由
	router.InitRouter(app)

	// logger
	log.InitLogger()
	app.Use(middleware.GinLogger(log.Logger), middleware.GinRecovery(log.Logger, true))
	defer func() {
		err := log.Logger.Sync()
		if err != nil {
			log.Logger.Error("log.Logger.Sync err", zap.Error(err))
		}
	}()
	err := app.Run(fmt.Sprintf("%s:%s", config.Conf.App.Host, config.Conf.App.Port))
	if err != nil {
		log.Logger.Error("main App Run error", zap.Error(err))
	}
}
