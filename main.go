package main

import (
	"fmt"
	"os"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/router"
	_ "github.com/yino/nlp-controller/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(fmt.Sprintf("%s:%s", config.Conf.App.Host, config.Conf.App.Port))
}
