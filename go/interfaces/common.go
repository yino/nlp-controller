package interfaces

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BaseResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type appReturn struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`

	Success        bool `json:"success"`
	HttpStatusCode int  `json:"httpStatusCode"`
}

func SendResp(c *gin.Context, args ...interface{}) {
	r := &appReturn{
		Code: StatusSuccess,
		Msg:  StatusText[StatusSuccess],
	}
	//配置下发日志每一条记录
	if c.Request.Method == "GET" {
		log.Println("[INFO][GET]request url-path is: %s, request parameter is: %v", c.Request.URL.Path, c.Request.URL.Query())
	} else if c.Request.Method == "POST" {
		log.Println("[INFO][POST]request url-path is: %s, request parameter is: %v", c.Request.URL.Path, c.Request.PostForm)
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			if v != StatusSuccess {
				r.Msg = "internal error"
				if msg, ok := StatusText[v]; ok {
					r.Msg = msg
				}
				r.Code = v
				log.Println("[ERROR] api interface error code: %d, msg: %s", r.Code, r.Msg)
			}
		case error:
			//pass
		case interface{}:
			r.Data = v
		}
	}
	c.IndentedJSON(http.StatusOK, r)

}
