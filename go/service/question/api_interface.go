package question

import (
	"fmt"
	"nlp/service/common"

	"github.com/gin-gonic/gin"
)

func init() {
	common.RegisterRouter(&common.ApiConfig{Path: "/question/index", Method: "get", Func: handleQuestionIndex})
}

func handleQuestionIndex(c *gin.Context) {
	c.JSON(200, "1312314asd")
}

func Test() {
	fmt.Println("test")
}
