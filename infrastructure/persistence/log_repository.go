package persistence

import (
	"github.com/yino/nlp-controller/domain/po"
	"gorm.io/gorm"
)

// QaQuestionRepo qa repo infra
type LogRepo struct {
	db *gorm.DB
}

// NewQaQuestionRepository qa test
func NewLogRepository(db *gorm.DB) *LogRepo {
	return &LogRepo{db: db}
}

// Add Add
func (log *LogRepo) Add(apiLog *po.APILog) error {
	return nil
}

//Page Page
func (log *LogRepo) Page(uid uint64, page, limit int64) (datList []po.APILog, total uint, err error) {
	return
}

// Count Count
func (log *LogRepo) Count(uid uint64) (count int64, err error) {
	return
}

// CountByAPIType CountByAPIType
func (log *LogRepo) CountByAPIType(uid uint64) (count int64, err error) {
	return
}

// CountByDay CountByDay
func (log *LogRepo) CountByDay(uid uint64, startTime, endTime int64) {

}
