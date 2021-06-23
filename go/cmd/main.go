package main

import (
	"nlp/service/common"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	common.ResisterGouterGin(r)
}
