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
		slaveQuestion.ID = question.ID
		qaPoSlaveList = append(qaPoSlaveList, slaveQuestion)
	}
	err := qa.domain.Edit(uid, qaMaster, qaPoSlaveList)
	if err != nil {
		return interfaces.ErrorUpdateData, err.Error()
	}

	return interfaces.StatusSuccess, ""
}

// Delete 删除
func (qa *QaQuestionApp) Delete(uid, id uint64) (int, string) {
	if err := qa.domain.Delete(uid, id); err != nil {
		return interfaces.ErrorDeleteData, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// Info 删除
func (qa *QaQuestionApp) Info(uid, id uint64) (int, vo.QaQuestionInfoVo) {
	infoVo, err := qa.domain.FindInfo(id)
	if err != nil {
		return interfaces.ErrorGetData, infoVo
	}

	if infoVo.UserId != uid {
		return interfaces.ErrorDataNoteUser, infoVo
	}

	return interfaces.StatusSuccess, infoVo
}

// Match 检索
func (qa *QaQuestionApp) Match(uid uint64, question string) (int, vo.QaMatchQuestionVo) {
	result, err := qa.domain.Match(uid, question)
	if err != nil {
		return interfaces.ErrorMatchQuestion, vo.QaMatchQuestionVo{}
	}

	return interfaces.StatusSuccess, vo.QaMatchQuestionVo{Data: result}
}

// Train 训练模型
func (qa *QaQuestionApp) Train(uid uint64) (int, string) {
	err := qa.domain.Train(uid)
	if err != nil {
		return interfaces.ErrorTrainQa, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// QuestionTotalNumber 获取问题总数
func (qa *QaQuestionApp) QuestionTotalNumber(uid uint64) (int, vo.QaQuestionTotal) {
	resp, err := qa.domain.QuestionTotalNumber(uid)
	if err != nil {
		return interfaces.ErrorQuestion, resp
	}
	return interfaces.StatusSuccess, resp
}

// NewQaApp new user app
func NewQaApp(qaRepo repository.QaQuestionRepository, userRepo repository.UserRepository) QaQuestionApp {
	return QaQuestionApp{
		domain: domain.NewQaDomain(qaRepo, userRepo),
	}
}
