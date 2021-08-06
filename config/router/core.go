package router

import (
	app "github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/infrastructure/persistence"
	"github.com/yino/nlp-controller/interfaces/corp"
	"github.com/yino/nlp-controller/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterCoreRouter 注册 core路由
func RegisterCoreRouter(c *gin.Engine, repo *persistence.Repositories) {
	userApp := app.NewUserApp(repo.User)
	userInterFace := corp.NewUsersInterface(userApp)

	qaInterFace := corp.NewQaInterface(app.NewQaApp(repo.Qa, repo.User))

	v1 := c.Group("v1")
	{
		core := v1.Group("core")
		{
			// 登录注册
			core.POST("/login", userInterFace.HandlerUserLogin)
			core.POST("/register", userInterFace.HandlerUserRegister)

			core.Use(middleware.CorpAuthTokenMiddleware(userApp))
			{
				// user
				core.GET("/user/info", userInterFace.HandlerUserInfo)
				core.GET("/user/edit", userInterFace.HandlerUserEdit)
				core.GET("/user/ak/page", userInterFace.HandlerUserAkPage)
				core.POST("/user/ak/add", userInterFace.HandlerUserCreateAk)
				core.GET("/user/ak/delete", userInterFace.HandlerUserAkDelete)

				// question
				core.GET("/question/index", qaInterFace.HandlerQuestionPage)
				core.POST("/question/add", qaInterFace.HandlerQuestionAdd)
				core.POST("/question/edit", qaInterFace.HandlerQuestionEdit)
				core.GET("/question/info", qaInterFace.HandlerQuestionInfo)
				core.GET("/question/delete", qaInterFace.HandlerQuestionDelete)
				core.GET("/question/train", qaInterFace.HandlerQuestionTrain)
				core.GET("/question/match", qaInterFace.HandlerQuestionMatch)
				core.GET("/question/total", qaInterFace.HandlerQuestionTotalNumber)
			}
		}
	}

}
