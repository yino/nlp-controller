package router

import (
	"github.com/gin-gonic/gin"
	"nlp/config/database"
	"nlp/infrastructure/persistence"
)

func InitRouter(c *gin.Engine) {

	_config := database.GetMysqlConf()
	repo, _ := persistence.NewRepositories(_config.MysqlUser, _config.MysqlPassword, _config.MysqlPort, _config.MysqlHost, _config.MysqlDb)
	RegisterCoreRouter(c, repo)
}
