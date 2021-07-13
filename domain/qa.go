package domain

import (
	"errors"

	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
)

// Qa qa question 领域服务
type Qa struct {
	QaRepo repository.QaQuestionRepository
}

// Page question page list
func (q *Qa) Page(page, limit int64, search map[string]interface{}) (list []po.QaQuestion, total int64, err error) {
	return q.QaRepo.Page(page, limit, search)
}

// Add add question
func (q *Qa) Add(qa *entity.QaQuestion) error {
	qaPo := new(po.QaQuestion)
	qaPo.Answer = qa.Answer
	qaPo.Question = qa.Question
	qaPo.Pid = qa.Pid
	qaPo.Type = qa.Type
	qaPo.UserId = qa.UserId
	return q.QaRepo.Add(qaPo)
}

// Edit edit qa
func (q *Qa) Edit(qa *entity.QaQuestion) error {
	qaPo, err := q.QaRepo.FindInfo(qa.ID)
	if err != nil {
		return err
	}
	qaPo.Answer = qa.Answer
	qaPo.Question = qa.Question
	qaPo.Pid = qa.Pid
	qaPo.Type = qa.Type
	return q.QaRepo.Edit(qaPo)
}

// Delete delete question
func (q *Qa) Delete(id uint64) error {
	return q.QaRepo.Delete(id)
}

// BatchInsert 批量插入数据
func (q *Qa) BatchInsert(data []entity.QaQuestion) error {

	var insertData []po.QaQuestion

	for _, qaEntity := range data {
		insertData = append(insertData, po.QaQuestion{
			Answer:   qaEntity.Answer,
			Question: qaEntity.Question,
			Pid:      qaEntity.Pid,
			Type:     qaEntity.Type,
			UserId:   qaEntity.UserId,
		})
	}
	return q.QaRepo.BatchInsert(insertData)
}

// FindInfo 根据id查询info
func (q *Qa) FindInfo(id uint64) (entity.QaQuestion, error) {
	data, err := q.QaRepo.FindInfo(id)
	var info entity.QaQuestion

	if data.ID == 0 {
		err = errors.New("data not found")
	}
	if err == nil {
		info.ID = data.ID
		info.Question = data.Question
		info.Answer = data.Answer
		info.Type = data.Type
	}

	return info, err
}

//NewQaDomain new qa domain
func NewQaDomain(repo repository.QaQuestionRepository) Qa {
	return Qa{
		QaRepo: repo,
	}
}
