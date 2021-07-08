package main

import (
	"fmt"
	"github.com/yino/nlp/config"
	"github.com/yino/nlp/config/router"
	"os"

	_ "github.com/yino/nlp/docs"

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
	// 注册路由
	router.InitRouter(app)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(fmt.Sprintf("%s:%s", config.Conf.App.Host, config.Conf.App.Port))
}
