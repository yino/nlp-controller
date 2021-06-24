package common

import (
	"github.com/gin-gonic/gin"
)

type ApiConfig struct {
	Path   string
	Method string
	Func   func(c *gin.Context)
}

var RouterList  = make([]*ApiConfig, 0, 10)

func RegisterRouter(register *ApiConfig) {
	RouterList = append(RouterList, register)
}

func ResisterRouterGin(r *gin.Engine) {

}
