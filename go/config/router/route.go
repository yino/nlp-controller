package router

import (
	"nlp/config"
	"nlp/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
}
