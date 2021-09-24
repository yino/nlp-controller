package persistence

import (
	"fmt"
	"time"

	"github.com/yino/nlp-controller/domain/po"
	"gorm.io/gorm"
)

// LogRepo .
type LogRepo struct {
	db *gorm.DB
}

// NewLogRepository .
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

// MaxQPS 统计最大的QPS
func (log *LogRepo) MaxQPS(uid uint64) (num int64, err error) {

	type Result struct {
		QPS int64 `json:"qps"`
	}
	var count int64
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Count(&count).Error
	if count == 0 || err != nil {
		return 0, nil
	}
	var res Result
	err = log.db.Model(&po.APILog{}).Where("user_id = ?", uid).Select("count(*) as qps").Group("created_at").Order("qps desc").First(&res).Error
	fmt.Println("err", err)
	if err == nil {
		num = res.QPS
	}
	return
}

// GroupCountBySecondOfDay GroupCountBySecondOfDay
func (log *LogRepo) GroupCountBySecondOfDay(uid uint64, startTime, endTime time.Time) (resp []po.APILogGroupTime, err error) {
	query := log.db.Model(&po.APILog{}).
		Where("user_id = ?", uid).
		Where("created_at between ? and ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).
		Select("DATE_FORMAT(date(created_at),'%Y-%m-%d') AS `datetime`")

	err = log.db.Table("(?) as a", query).
		Debug().
		Select("count(*) as total, datetime").
		Group("DATE_FORMAT(`datetime`, '%Y-%m-%d')").
		Order("datetime").
		Scan(&resp).
		Error
	return
}

// CountByDay 分组统计
func (log *LogRepo) CountByDay(uid uint64, startTime, endTime time.Time) (resp []po.APILogGroupTime, err error) {
	query := log.db.Model(&po.APILog{}).
		Where("user_id = ?", uid).
		Where("created_at between ? and ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).
		Select("DATE_FORMAT(date(created_at),'%Y-%m-%d') AS `datetime`")

	err = log.db.Table("(?) as a", query).
		Debug().
		Select("count(*) as total, datetime").
		Group("DATE_FORMAT(`datetime`, '%Y-%m-%d')").
		Order("datetime").
		Scan(&resp).
		Error
	return
}

// CountByDayByAPIStatus 分组统计 并按 api status 筛选
func (log *LogRepo) CountByDayByAPIStatus(uid uint64, startTime, endTime time.Time, apiStatus string) (resp []po.APILogGroupTime, err error) {
	query := log.db.Model(&po.APILog{}).
		Where("user_id = ?", uid).
		Where("created_at between ? and ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).
		Where("api_status = ?", apiStatus).
		Select("DATE_FORMAT(date(created_at),'%Y-%m-%d') AS `datetime`")

	err = log.db.Table("(?) as a", query).
		Debug().
		Select("count(*) as total, datetime").
		Group("DATE_FORMAT(`datetime`, '%Y-%m-%d')").
		Order("datetime").
		Scan(&resp).
		Error
	return
}
