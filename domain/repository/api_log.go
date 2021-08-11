package repository

import (
	"time"

	"github.com/yino/nlp-controller/domain/po"
)

// APILogRepository api
type APILogRepository interface {
	Add(log *po.APILog) error
	Page(uid uint64, page, limit int64) (datList []po.APILog, total uint, err error)
	Count(uid uint64) (int64, error)
	CountByAPIType(uid uint64, apiType string) (int64, error)
	// CountByDay 按天统计
	CountByDay(uid uint64, startTime, endTime int64)
	CountByNormalStatus(uid uint64, status string) (int64, error)
	// GroupCountBySecond 按秒分组统计当天
	GroupCountBySecondOfDay(uid uint64, startTime, endTime time.Time, limit int64) ([]po.APILogGroupTime, error)
}
