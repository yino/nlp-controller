package application

import (
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"
)

type LogApp struct {
	domain domain.Log
}

func (l *LogApp) QPS(uid uint64, startTime, endTime int64) ([]vo.LogQPS, int) {
	resp, err := l.domain.QPS(uid, startTime, endTime)
	if err != nil {
		return nil, interfaces.ErrorLogQPS
	}
	return resp, interfaces.StatusSuccess
}

func (l *LogApp) Write(uid uint64, method string, params []byte, header []byte, ip, URL, apiStatus string) int {
	logEntity := new(entity.Log)
	logEntity.APILog = po.APILog{
		Method:    method,
		Params:    params,
		IP:        ip,
		Header:    header,
		UserID:    uid,
		APIType:   domain.QaType,
		APIStatus: apiStatus,
		URL:       URL,
	}
	err := l.domain.Add(logEntity)
	if err != nil {
		return interfaces.ErrorLogQPS
	}
	return interfaces.StatusSuccess
}

// NewLogApp new user app
func NewLogApp(repo repository.APILogRepository) LogApp {
	return LogApp{
		domain: domain.NewLogDomain(repo),
	}
}
