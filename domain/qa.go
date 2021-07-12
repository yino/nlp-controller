package domain

import (
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
)

// QaDomain qa question 领域服务
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
	qaPo := new(po.QaQuestion)
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
func (q *Qa) BatchInsert(data []po.QaQuestion) error {
	return q.QaRepo.BatchInsert(data)
}

// FindInfo 根据id查询info
func (q *Qa) FindInfo(id uint64) error {
	_, err := q.QaRepo.FindInfo(id)
	return err
}

func NewQaDomain(repo repository.QaQuestionRepository) Qa {
	return Qa{
		QaRepo: repo,
	}
}
