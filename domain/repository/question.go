package repository

import (
	"github.com/yino/nlp-controller/domain/po"
)

// QaQuestionRepository: QA问题库聚合工厂
type QaQuestionRepository interface {
	Page(page, limit int64, search map[string]string) ([]po.QaQuestion, error)
	Add(question *po.QaQuestion) error
	Edit(question *po.QaQuestion) error
	Delete(id uint64) error
	BatchInsert([]po.QaQuestion) error
}
