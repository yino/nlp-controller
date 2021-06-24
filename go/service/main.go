package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"nlp/service/common"
	"os"
	"strings"
)
var apis []*common.ApiConfig
func init()  {
	apis = common.RouterList
}

func main() {
	os.Setenv("GIN_MODE", "release")

	router := gin.New()
	//router.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true), gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", func(context *gin.Context) {
		context.JSON(200,"123123")
	})
	v1 := router.Group("v1")
	{
		fmt.Println(apis)
		if apis != nil {
			for _, apiConf := range apis {
				switch strings.ToLower(apiConf.Method) {
				case "post":
					v1.POST(apiConf.Path, apiConf.Func)
				case "get":
					v1.GET(apiConf.Path, apiConf.Func)
				case "delete":
					v1.DELETE(apiConf.Path, apiConf.Func)
				case "patch":
					v1.PATCH(apiConf.Path, apiConf.Func)
				case "put":
					v1.PUT(apiConf.Path, apiConf.Func)
				case "options":
					v1.OPTIONS(apiConf.Path, apiConf.Func)
				case "head":
					v1.HEAD(apiConf.Path, apiConf.Func)
				default:
					log.Fatalf("|debug|unsuported method|%s|%s|", apiConf.Method, apiConf.Path)
				}
			}
		}
	}

	router.Run("localhost:8080")
}
