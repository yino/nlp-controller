package middleware

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/interfaces"
)

// APIAkAuthMiddleware auth token middleware
func APIAkAuthMiddleware(user application.UserApp, logApp application.LogApp) gin.HandlerFunc {

	return func(c *gin.Context) {
		ak := c.Query("ak")
		if len(ak) == 0 {
			c.Abort()
			interfaces.SendResp(c, interfaces.ErrorNotLogin)
			return
		}
		vo, ret := user.FindAkByUser(ak)
		if ret != interfaces.StatusSuccess {
			c.Abort()
			interfaces.SendResp(c, ret)
			return
		}
		c.Set("uid", vo.Id)
		c.Next()
		params, _ := ioutil.ReadAll(c.Request.Body)
		params = []byte("")
		logApp.Write(vo.Id, c.Request.Method, params, params, c.Request.Host, c.Request.RequestURI)
	}

}
