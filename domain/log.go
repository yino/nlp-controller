package domain

import (
	"fmt"
	"time"

	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
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
func (l *Log) QPS(uid uint64, startTime, endTime int64) (resp []vo.LogQPS, err error) {
	var beginTime, OffTime time.Time
	beginTime = time.Unix(startTime, 0)
	OffTime = time.Unix(endTime, 0)
	result, err := l.APILogRepo.GroupCountBySecondOfDay(uid, beginTime, OffTime)
	if err != nil {
		return
	}
	datetimeMap := make(map[string]int64)
	for _, val := range result {
		datetimeMap[val.Datetime] = val.Total
	}

	for i := startTime; i <= endTime; i++ {
		total := int64(0)
		dateStr := time.Unix(i, 0).Format("2006-01-02 15:04:05")
		if v, ok := datetimeMap[dateStr]; ok {
			total = v
		}
		resp = append(resp, vo.LogQPS{
			Datetime: dateStr,
			Total:    total,
		})
	}
	return

}

// NewUserDomain new domain.Log
func NewLogDomain(repo repository.APILogRepository) Log {
	return Log{
		APILogRepo: repo,
	}
}
