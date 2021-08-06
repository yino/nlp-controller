package router

import (
	"net/http"

	"github.com/yino/nlp-controller/config"
	_ "github.com/yino/nlp-controller/docs"
	"github.com/yino/nlp-controller/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(c *gin.Engine) {
	c.Use(Cors())
	c.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 跨域解决

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	//repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
	RegisterAPIRouter(c, repo)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 允许 Origin 字段中的域发送请求
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// OPTIONS请求返回200
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
