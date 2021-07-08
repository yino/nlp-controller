package router

import (
	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/infrastructure/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	c.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
}
