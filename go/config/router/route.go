package router

import (
	"github.com/gin-gonic/gin"
	"nlp/config"
	"nlp/infrastructure/persistence"
)

func InitRouter(c *gin.Engine) {

	//_config := database.GetMysqlConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	repo.AutoMigrate()
	RegisterCoreRouter(c, repo)
}
