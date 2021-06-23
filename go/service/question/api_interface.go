package question

import (
	"nlp/service/common"

	"github.com/gin-gonic/gin"
)

func init() {
	common.RegisterRouter(common.RouteData{"/question/index", "get", handleQuestionIndex})
}

func handleQuestionIndex(c *gin.Context) {

}
