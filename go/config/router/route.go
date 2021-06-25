package router

import "github.com/gin-gonic/gin"

func InitRouter(c *gin.Engine){
	RegisterCoreRouter(c)
}
