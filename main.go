package main

import (
	"fmt"
	"os"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/config/router"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	config.GetConf()
	os.Setenv("GIN_MODE", "debug")
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	pprof.Register(app)
	// 注册路由
	router.InitRouter(app)

	// logger
	log.InitLogger()
	defer log.Logger.Sync()
	app.Run(fmt.Sprintf("%s:%s", config.Conf.App.Host, config.Conf.App.Port))
}
