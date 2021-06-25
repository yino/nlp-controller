package router

import (
	"github.com/gin-gonic/gin"
	"nlp/interfaces"
)

func RegisterCoreRouter(c *gin.Engine) {

	v1 := c.Group("v1")
	{
		core := v1.Group("core")
		{
			// question
			core.GET("/question/index", interfaces.HandlerQuestionPage)
			core.POST("/question/add", interfaces.HandlerQuestionAdd)
			core.POST("/question/edit", interfaces.HandlerQuestionEdit)
			core.GET("/question/delete", interfaces.HandlerQuestionDelete)
			core.GET("/question/train", interfaces.HandlerQuestionTrain)

			// user
			core.GET("/user/info")
			core.GET("/user/edit")

			// login
			core.POST("/login", interfaces.HandlerLogin)
		}
	}

}
