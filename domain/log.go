package domain

import (
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
)

// Log log domain
type Log struct {
	APILogRepo repository.APILogRepository // api log 聚合工厂
}

// Add 插入数据
func (l *Log) Add(data *entity.Log) error {
	apiLogPo := data.APILog
	return l.APILogRepo.Add(&apiLogPo)
}

// NewUserDomain new domain.Log
func NewLogDomain(repo repository.APILogRepository) Log {
	return Log{
		APILogRepo: repo,
	}
}
