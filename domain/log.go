package domain

import (
	"fmt"
	"time"

	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
)

// Log log domain
type Log struct {
	APILogRepo repository.APILogRepository // api log 聚合工厂
}

// Add 插入数据
func (l *Log) Add(data *entity.Log) error {
	apiLogPo := &data.APILog
	fmt.Println(apiLogPo)
	return l.APILogRepo.Add(apiLogPo)
}

// RequestTotalNum 请求量
func (l *Log) RequestTotalNum(uid uint64, status string) (total int64, err error) {
	if status == ALL {
		return l.APILogRepo.Count(uid)
	}
	return l.APILogRepo.CountByNormalStatus(uid, status)
}

// QPS QPS
func (l *Log) QPS(uid uint64, startTime, endTime, limit int64) {
	var beginTime, OffTime time.Time
	returl.APILogRepo.GroupCountBySecondOfDay(uid, beginTime, OffTime, limit)
}

// NewUserDomain new domain.Log
func NewLogDomain(repo repository.APILogRepository) Log {
	return Log{
		APILogRepo: repo,
	}
}
