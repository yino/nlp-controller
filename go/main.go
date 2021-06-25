package main

import (
	"github.com/gin-gonic/gin"
	"nlp/config/router"
	"os"
)

func main() {
	os.Setenv("GIN_MODE", "release")
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	// 注册路由
	router.InitRouter(app)

	app.Run("localhost:8080")
}
