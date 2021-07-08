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
	c.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
}
