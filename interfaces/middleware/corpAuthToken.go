package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/interfaces"
)

// CorpAuthTokenMiddleware auth token middleware
func CorpAuthTokenMiddleware(user application.UserApp) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.Abort()
			interfaces.SendResp(c, interfaces.ErrorNotLogin)
			return
		}
		vo, ret := user.AuthToken(token)
		if ret != interfaces.StatusSuccess {
			c.Abort()
			fmt.Println("2222")
			interfaces.SendResp(c, ret)
			return
		}
		c.Set("uid", vo.Id)
		c.Next()
	}

}
