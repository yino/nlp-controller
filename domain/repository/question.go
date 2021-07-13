package repository

import (
	"github.com/yino/nlp-controller/domain/po"
)

// QaQuestionRepository QA问题库聚合工厂
type QaQuestionRepository interface {
	Page(page, limit int64, search map[string]interface{}) (list []po.QaQuestion, total int64, err error)
	AddMaster(question *po.QaQuestion) (uint64, error)
	EditMaster(question *po.QaQuestion) error
	Delete(id uint64) error
	BatchInsert([]po.QaQuestion) error
	FindInfo(id uint64) (*po.QaQuestion, error)
	GetSlaveList(pid uint64) ([]po.QaQuestion, error)
	Add(masterQuestion *po.QaQuestion, slave []po.QaQuestion) error
	Edit(masterQuestion *po.QaQuestion, slave []po.QaQuestion) error
}
