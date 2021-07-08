package router

import (
	"github.com/yino/nlp/config"
	"github.com/yino/nlp/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
}
