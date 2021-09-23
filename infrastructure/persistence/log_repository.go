package persistence

import (
	"fmt"
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

// CountByNormalStatus 根据uid与api status 获取请求量
func (log *LogRepo) CountByNormalStatus(uid uint64, status string) (count int64, err error) {
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Where("api_status", status).Count(&count).Error
	return
}

// CountByDay 统计最大的QPS
func (log *LogRepo) MaxQPS(uid uint64) (num int64, err error) {

	type Result struct {
		QPS int64 `json:"qps"`
	}
	var res Result
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Select("count(*) as qps").Group("created_at").Order("qps desc").First(&res).Error

	if err == nil {
		num = res.QPS
	}
	return
}

// GroupCountBySecondOfDay GroupCountBySecondOfDay
func (log *LogRepo) GroupCountBySecondOfDay(uid uint64, startTime, endTime time.Time) (resp []po.APILogGroupTime, err error) {
	where := fmt.Sprintf("where user_id=%d and created_at between '%s' and '%s'", uid, startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
	sqlStr := "SELECT `datetime`, COUNT( * ) AS total FROM(SELECT *,DATE_FORMAT(concat( date( created_at ), ' ',HOUR ( created_at ), ':', MINUTE ( created_at ),':', SECOND(created_at)),'%Y-%m-%d %H:%i:%s') AS `datetime` FROM api_log " + where + ") a GROUP BY DATE_FORMAT( `datetime`, '%Y-%m-%d %H:%i' ) ORDER BY `datetime`"
	err = log.db.Raw(sqlStr).Scan(&resp).Error
	return
}
