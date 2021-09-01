package corp

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/interfaces"
)

type APILog struct {
	logApp application.LogApp
}

func (a *APILog) QPS(c *gin.Context) {
	uid := GetUid(c)

	startTime, err := strconv.ParseInt(c.Query("startTime"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
		return
	}

	endTime, err := strconv.ParseInt(c.Query("endTime"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
	}

	if endTime-startTime > 3600*24 {
		interfaces.SendResp(c, interfaces.ErrorParams)
	}

	resp, ret := a.logApp.QPS(uid, startTime, endTime)
	interfaces.SendResp(c, ret, resp)
}

// NewAPILogInterface new APILog interface
func NewAPILogInterface(app application.LogApp) APILog {
	return APILog{
		logApp: app,
	}
}
