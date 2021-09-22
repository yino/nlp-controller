package router

import (
	app "github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/infrastructure/persistence"
	"github.com/yino/nlp-controller/interfaces/corp"
	"github.com/yino/nlp-controller/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRouter api 路由
func RegisterAPIRouter(c *gin.Engine, repo *persistence.Repositories) {
	userApp := app.NewUserApp(repo)
	userInterFace := corp.NewUsersInterface(userApp)

	qaInterFace := corp.NewQaInterface(app.NewQaApp(repo))
	logApp := app.NewLogApp(repo)
	v1 := c.Group("v1")
	{
		core := v1.Group("api")
		{
			core.Use(middleware.APIAkAuthMiddleware(userApp, logApp))
			{
				// user
				core.GET("/user/info", userInterFace.HandlerUserInfo)

				// question
				core.GET("/question/index", qaInterFace.HandlerQuestionPage)
				core.POST("/question/add", qaInterFace.HandlerQuestionAdd)
				core.POST("/question/edit", qaInterFace.HandlerQuestionEdit)
				core.GET("/question/info", qaInterFace.HandlerQuestionInfo)
				core.GET("/question/delete", qaInterFace.HandlerQuestionDelete)
				core.GET("/question/train", qaInterFace.HandlerQuestionTrain)
				core.GET("/question/match", qaInterFace.HandlerQuestionMatch)
			}
		}
	}

}
