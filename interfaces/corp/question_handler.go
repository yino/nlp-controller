package corp

import (
	"strconv"

	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"

	"github.com/gin-gonic/gin"
)

// Qa domain
type Qa struct {
	qa application.QaQuestionApp
}

// HandlerQuestionPage page
func (q *Qa) HandlerQuestionPage(c *gin.Context) {
	uid := GetUid(c)
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
	}
	vo, err := q.qa.Page(uid, page, limit)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorGetData)
	}

	interfaces.SendResp(c, interfaces.StatusSuccess, vo)
}

// HandlerQuestionAdd add
func (q *Qa) HandlerQuestionAdd(c *gin.Context) {
	var req vo.QaAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	ret, msg := q.qa.Add(GetUid(c), req)
	interfaces.SendResp(c, ret, msg)
}

// HandlerQuestionEdit edit
func (q *Qa) HandlerQuestionEdit(c *gin.Context) {
	var req vo.QaEditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	ret, msg := q.qa.Edit(GetUid(c), req)
	interfaces.SendResp(c, ret, msg)
}

// HandlerQuestionDelete delete
func (q *Qa) HandlerQuestionDelete(c *gin.Context) {
	reqID := c.Query("id")

	id, err := strconv.ParseUint(reqID, 10, 64)
	if err != nil || id == 0 {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	ret, msg := q.qa.Delete(GetUid(c), id)
	interfaces.SendResp(c, ret, msg)
}

// HandlerQuestionInfo get info
func (q *Qa) HandlerQuestionInfo(c *gin.Context) {
	reqID := c.Query("id")

	id, err := strconv.ParseUint(reqID, 10, 64)
	if err != nil || id == 0 {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	ret, data := q.qa.Info(GetUid(c), id)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret)
		return

	}
	interfaces.SendResp(c, ret, data)
}

// HandlerQuestionTrain training qa model
func (q *Qa) HandlerQuestionTrain(c *gin.Context) {

	ret, data := q.qa.Train(GetUid(c))
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, data)
		return
	}
	interfaces.SendResp(c, ret, data)
}

// HandlerQuestionMatch model match question
func (q *Qa) HandlerQuestionMatch(c *gin.Context) {
	question := c.DefaultQuery("question", "")
	if len(question) == 0 {
		interfaces.SendResp(c, interfaces.ErrorParams)
		return
	}
	ret, data := q.qa.Match(GetUid(c), question)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, data)
		return
	}
	interfaces.SendResp(c, ret, data)
}

// HandlerQuestionTotalNumber get total number
func (q *Qa) HandlerQuestionTotalNumber(c *gin.Context) {
	ret, data := q.qa.QuestionTotalNumber(GetUid(c))
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, data)
		return
	}
	interfaces.SendResp(c, ret, data)
}

// NewQaInterface new qa interface
func NewQaInterface(app application.QaQuestionApp) Qa {
	return Qa{
		qa: app,
	}
}
