package persistence

import (
	"time"

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
	return log.db.Create(apiLog).Error
}

//Page Page
func (log *LogRepo) Page(uid uint64, page, limit int64) (datList []po.APILog, total uint, err error) {
	return
}

// Count Count
func (log *LogRepo) Count(uid uint64) (count int64, err error) {
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Count(&count).Error
	return
}

// CountByAPIType CountByAPIType
func (log *LogRepo) CountByAPIType(uid uint64, apiType string) (count int64, err error) {
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Where("api_type", apiType).Count(&count).Error
	return
}

// CountByNormalStatus CountByNormalStatus
func (log *LogRepo) CountByNormalStatus(uid uint64, status string) (count int64, err error) {
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Where("api_status", status).Count(&count).Error
	return
}

// CountByDay CountByDay
func (log *LogRepo) CountByDay(uid uint64, startTime, endTime int64) {

}

// GroupCountBySecondOfDay GroupCountBySecondOfDay
func (log *LogRepo) GroupCountBySecondOfDay(uid uint64, startTime, endTime time.Time, limit int64) (result []po.APILogGroupTime, err error) {
	var resp []po.APILogGroupTime
	err = log.db.Raw("SELECT `time`, COUNT( * ) AS total FROM(SELECT *,DATE_FORMAT(concat( date( created_at ), ' ',HOUR ( created_at ), ':', MINUTE ( created_at ),':', SECOND(created_at)),'%Y-%m-%d %H:%i:%s') AS `time` FROM api_log) a GROUP BY DATE_FORMAT( `time`, '%Y-%m-%d %H:%i' ) ORDER BY `time`").Scan(&resp).Error
	return
}
