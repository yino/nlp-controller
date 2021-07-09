package persistence

import (
	"github.com/yino/nlp-controller/domain/po"

	"gorm.io/gorm"
)

type QaQuestionRepo struct {
	db *gorm.DB
}

// NewUserRepository： user test
func NewQaQuestionRepository(db *gorm.DB) *QaQuestionRepo {
	return &QaQuestionRepo{db: db}
}
func (repo *QaQuestionRepo) Page(page, limit int64, search map[string]string) ([]po.QaQuestion, error) {
	return []po.QaQuestion{}, nil
}

func (repo *QaQuestionRepo) Add(*po.QaQuestion) error {
	return nil
}

func (repo *QaQuestionRepo) Edit(*po.QaQuestion) error {
	return nil
}

func (repo *QaQuestionRepo) Delete(id uint64) error {
	return nil
}

func (repo *QaQuestionRepo) BatchInsert([]po.QaQuestion) error {
	return nil
}
