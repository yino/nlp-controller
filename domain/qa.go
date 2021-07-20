package domain

import (
	"errors"

	"github.com/yino/nlp-controller/domain/vo"

	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
)

// Qa qa question 领域服务
type Qa struct {
	QaRepo repository.QaQuestionRepository
}

// GetMasterQuestionPage 获取 master
func (q *Qa) GetMasterQuestionPage(page, limit int64, search map[string]interface{}) (list vo.QaQuestionPageVo, err error) {
	search["pid"] = 0
	dataList, total, err := q.QaRepo.Page(page, limit, search)
	if err == nil && len(dataList) > 0 {
		for _, qaInfo := range dataList {
			list.Data = append(list.Data, vo.QaQuestionVo{
				Id:       qaInfo.ID,
				Question: qaInfo.Question,
				Answer:   qaInfo.Answer,
			})
		}
	}
	list.Page = page
	list.Total = total
	list.PageSize = limit
	return
}

// AddMaster add question
func (q *Qa) AddMaster(qa *entity.QaQuestion) error {
	qaPo := new(po.QaQuestion)
	qaPo.Answer = qa.Answer
	qaPo.Question = qa.Question
	qaPo.Pid = qa.Pid
	qaPo.Type = qa.Type
	qaPo.UserId = qa.UserId
	_, err := q.QaRepo.AddMaster(qaPo)
	return err
}

// EditMaster edit qa
func (q *Qa) EditMaster(qa *entity.QaQuestion) error {
	qaPo, err := q.QaRepo.FindInfo(qa.ID)
	if err != nil {
		return err
	}
	qaPo.Answer = qa.Answer
	qaPo.Question = qa.Question
	qaPo.Pid = qa.Pid
	qaPo.Type = qa.Type
	return q.QaRepo.EditMaster(qaPo)
}

// Delete delete question
func (q *Qa) Delete(uid, id uint64) error {
	qaPo, err := q.QaRepo.FindInfo(id)
	if qaPo.ID == 0 {
		return errors.New("data not found")
	}
	if err != nil {
		return err
	}
	return q.QaRepo.Delete(id)
}

// FindInfo 根据id查询info
func (q *Qa) FindInfo(id uint64) (vo.QaQuestionInfoVo, error) {
	data, err := q.QaRepo.FindInfo(id)
	var info vo.QaQuestionInfoVo

	if data.ID == 0 {
		err = errors.New("data not found")
	}
	if err != nil {
		return info, err
	}
	info.Id = data.ID
	info.Question = data.Question
	info.Answer = data.Answer
	info.UserId = data.UserId

	slaveQuestionList, _ := q.QaRepo.GetSlaveList(data.ID)
	if err == nil && len(slaveQuestionList) > 0 {
		for _, slaveQuestion := range slaveQuestionList {
			info.SimilarQuestion = append(info.SimilarQuestion, vo.QaQuestionVo{
				Id:       slaveQuestion.ID,
				Question: slaveQuestion.Question,
				Answer:   slaveQuestion.Answer,
			})
		}
	}
	return info, err
}

// Add add master question and slave questions
func (q *Qa) Add(masterQuestion *entity.QaQuestion, slaveQuestion []entity.QaQuestion) error {
	qaPo := new(po.QaQuestion)
	qaPo.Answer = masterQuestion.Answer
	qaPo.Question = masterQuestion.Question
	qaPo.Pid = 0
	qaPo.Type = 1
	qaPo.UserId = masterQuestion.UserId

	var qaPoSlaveList []po.QaQuestion
	for _, question := range slaveQuestion {
		qaPoSlaveList = append(qaPoSlaveList, po.QaQuestion{
			Question: question.Question,
			Answer:   question.Answer,
			Type:     2,
			UserId:   question.UserId,
			Pid:      0,
		})
	}
	return q.QaRepo.Add(qaPo, qaPoSlaveList)
}

// Edit master question and slave questions
func (q *Qa) Edit(uid uint64, masterQuestion *entity.QaQuestion, slaveQuestion []entity.QaQuestion) error {
	qaPo, err := q.QaRepo.FindInfo(masterQuestion.ID)
	if qaPo.ID == 0 {
		return errors.New("data not found")
	}
	if err != nil {
		return err
	}
	if qaPo.UserId != uid {
		return errors.New("not permission")
	}
	qaPo.Answer = masterQuestion.Answer
	qaPo.Question = masterQuestion.Question
	var qaPoSlaveList []po.QaQuestion
	for _, question := range slaveQuestion {
		qaPoSlaveList = append(qaPoSlaveList, po.QaQuestion{
			Question: question.Question,
			Answer:   question.Answer,
			Type:     2,
			UserId:   qaPo.UserId,
			Pid:      qaPo.ID,
			ID:       question.ID,
		})
	}
	return q.QaRepo.Edit(qaPo, qaPoSlaveList)
}

//NewQaDomain new qa domain
func NewQaDomain(repo repository.QaQuestionRepository) Qa {
	return Qa{
		QaRepo: repo,
	}
}
