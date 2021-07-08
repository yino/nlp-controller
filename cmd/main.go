package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nlp/service/common"
	"os"
)

func main() {
	os.Setenv("GIN_MODE", "release")

	router := gin.New()
	//router.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true), gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	common.ResisterRouterGin(router)

	routerList := common.RouterList
	fmt.Println(routerList)
}
