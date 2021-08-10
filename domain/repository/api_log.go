package repository

import "github.com/yino/nlp-controller/domain/po"

// APILogRepository api
type APILogRepository interface {
	Add(log *po.APILog) error
	Page(uid uint64, page, limit int64) (datList []po.APILog, total uint, err error)
	Count(uid uint64) (int64, error)
	CountByAPIType(uid uint64) (int64, error)
	// CountByDay 按天统计
	CountByDay(uid uint64, startTime, endTime int64)
}
