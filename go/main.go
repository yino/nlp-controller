package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nlp/config"
	"nlp/config/router"
)

func main() {
	// load config
	config.GetConf()
	//os.Setenv("GIN_MODE", "debug")
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	// 注册路由
	router.InitRouter(app)

	app.Run(fmt.Sprintf("%s:%s", config.Conf.App.Host, config.Conf.App.Port))
}
