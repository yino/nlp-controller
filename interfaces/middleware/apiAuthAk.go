package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

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

		// params
		var (
			params    []byte
			paramsMap = make(map[string]interface{})
			headerMap = make(map[string]interface{})
		)

		if c.Request.Method == "GET" {
			paramsArrMap, _ := url.ParseQuery(c.Request.URL.RawQuery)
			for k, v := range paramsArrMap {
				if len(v) > 1 {
					paramsMap[k] = v
				} else {
					paramsMap[k] = v[0]
				}
			}
			params, _ = json.Marshal(paramsMap)
		} else if c.Request.Method == "POST" {
			params, _ = ioutil.ReadAll(c.Request.Body)
		}

		// header
		for k, v := range c.Request.Header {
			if len(v) > 1 {
				headerMap[k] = v
			} else {
				headerMap[k] = v[0]
			}
		}
		header, _ := json.Marshal(headerMap)

		logApp.Write(vo.Id, c.Request.Method, params, header, c.Request.Host, c.Request.RequestURI)
	}

}
