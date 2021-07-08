package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	app "github.com/yino/nlp/application"
	"github.com/yino/nlp/infrastructure/persistence"
	"github.com/yino/nlp/interfaces/corp"
)

func RegisterCoreRouter(c *gin.Engine, repo *persistence.Repositories) {
	userApp := app.NewUserApp(repo.User)
	UserInterFace := corp.NewUsersInterface(userApp)
	fmt.Println("user interface")
	v1 := c.Group("v1")
	{
		core := v1.Group("core")
		{
			// question
			//core.GET("/question/index", interfaces.HandlerQuestionPage)
			//core.POST("/question/add", interfaces.HandlerQuestionAdd)
			//core.POST("/question/edit", interfaces.HandlerQuestionEdit)
			//core.GET("/question/delete", interfaces.HandlerQuestionDelete)
			//core.GET("/question/train", interfaces.HandlerQuestionTrain)

			// user
			core.GET("/user/info", UserInterFace.HandlerUserInfo)
			core.GET("/user/edit", UserInterFace.HandlerUserEdit)
			// 登录注册
			core.POST("/login", UserInterFace.HandlerUserLogin)
			core.POST("/register", UserInterFace.HandlerUserRegister)
		}
	}

}
