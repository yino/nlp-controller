package application

import (
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"
)

// QaQuestionApp qa app
type QaQuestionApp struct {
	//questionDomain
	domain domain.Qa
}

// Page 分页
func (qa *QaQuestionApp) Page(uid uint64, page, limit int64) (vo.QaQuestionPageVo, error) {
	search := make(map[string]interface{})
	search["user_id"] = uid
	return qa.domain.GetMasterQuestionPage(page, limit, search)
}

// Add 新增
func (qa *QaQuestionApp) Add(uid uint64, req vo.QaAddReq) (int, string) {
	qaMaster := new(entity.QaQuestion)
	qaMaster.Question = req.Question
	qaMaster.Answer = req.Answer
	qaMaster.UserId = uid
	var qaPoSlaveList []entity.QaQuestion
	for _, question := range req.SlaveQuestion {
		var slaveQuestion entity.QaQuestion
		slaveQuestion.UserId = uid
		slaveQuestion.Question = question
		slaveQuestion.Answer = req.Answer
		qaPoSlaveList = append(qaPoSlaveList, slaveQuestion)
	}
	err := qa.domain.Add(qaMaster, qaPoSlaveList)
	if err != nil {
		return interfaces.ErrorCreateData, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// Edit 编辑
func (qa *QaQuestionApp) Edit(uid uint64, req vo.QaEditReq) (int, string) {

	infoVo, err := qa.domain.FindInfo(req.ID)
	if err != nil {
		return interfaces.ErrorGetData, err.Error()
	}

	if infoVo.UserId != uid {
		return interfaces.ErrorDataNoteUser, err.Error()
	}

	qaMaster := new(entity.QaQuestion)
	qaMaster.Question = req.Question
	qaMaster.Answer = req.Answer
	qaMaster.ID = req.ID
	var qaPoSlaveList []entity.QaQuestion

	for _, question := range req.SlaveQuestion {
		var slaveQuestion entity.QaQuestion
		slaveQuestion.UserId = uid
		slaveQuestion.Question = question.Question
		slaveQuestion.Answer = req.Answer
		qaPoSlaveList = append(qaPoSlaveList, slaveQuestion)
	}
	err = qa.domain.Edit(qaMaster, qaPoSlaveList)
	if err != nil {
		return interfaces.ErrorUpdateData, err.Error()
	}

	return interfaces.StatusSuccess, ""
}

// Delete 删除
func (qa *QaQuestionApp) Delete(uid, id uint64) (int, string) {
	infoVo, err := qa.domain.FindInfo(id)
	if err != nil {
		return interfaces.ErrorGetData, err.Error()
	}

	if infoVo.UserId != uid {
		return interfaces.ErrorDataNoteUser, err.Error()
	}
	if err := qa.domain.Delete(id); err != nil {
		return interfaces.ErrorDeleteData, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// Match 检索
func (qa *QaQuestionApp) Match(uid, page, limit int64) {

}

// NewQaApp new user app
func NewQaApp(repo repository.QaQuestionRepository) QaQuestionApp {
	return QaQuestionApp{
		domain: domain.NewQaDomain(repo),
	}
}
