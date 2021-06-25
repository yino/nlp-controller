package router

import (
	"github.com/gin-gonic/gin"
	"nlp/service/question"
)

func RegisterCoreRouter(c *gin.Engine) {

	v1 := c.Group("v1")
	{
		core := v1.Group("corp")
		{
			// question
			core.GET("/question/index", question.HandleQuestionIndex)
			core.POST("/question/add", question.HandleQuestionAdd)
			core.POST("/question/edit", question.HandleQuestionEdit)
			core.GET("/question/delete", question.HandleQuestionDel)
			core.GET("/question/train", question.HandleQuestionTrain)

			// user
			core.GET("/user/info")
			core.GET("/user/edit")
		}
	}

}
