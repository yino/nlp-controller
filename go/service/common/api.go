package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RouteData struct {
	Path   string
	Method string
	Func   func(c *gin.Context)
}

var RouterList []RouteData

func RegisterRouter(register RouteData) {
	RouterList = append(RouterList, register)
}

func ResisterGouterGin(r *gin.Engine) {
	for index, routerData := range RouterList {
		fmt.Println(index, routerData)
	}
}
