package application

import (
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/infrastructure/persistence"
	"github.com/yino/nlp-controller/interfaces"
	"go.uber.org/zap"
)

// LogApp .
type LogApp struct {
	domain domain.Log
}

// QPS .
func (l *LogApp) QPS(uid uint64, startTime, endTime int64) ([]vo.LogQPS, int) {
	resp, err := l.domain.QPS(uid, startTime, endTime)
	if err != nil {
		return nil, interfaces.ErrorLogQPS
	}
	return resp, interfaces.StatusSuccess
}

// Write 写入log
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
		log.Logger.Error("get RequestNum validTotal", zap.Error(err))
		return interfaces.ErrorLogQPS
	}
	return interfaces.StatusSuccess
}

// RequestNum 统计请求量
func (l *LogApp) RequestNum(uid uint64) (vo.RequestNum, int) {
	var (
		rep vo.RequestNum
		ret = interfaces.StatusSuccess
	)
	// 无效的请求
	requestTotal, err := l.domain.RequestTotalNum(uid, domain.ALL)
	if err != nil {
		requestTotal = 0
		ret = interfaces.ErrorRequestNum
		log.Logger.Error("get RequestNum requestTotal", zap.Error(err))
	}
	// 有效的请求
	validTotal, err := l.domain.RequestTotalNum(uid, domain.NORMAL)
	if err != nil {
		validTotal = 0
		ret = interfaces.ErrorRequestNum
		log.Logger.Error("get RequestNum validTotal", zap.Error(err))
	}
	rep.RequestTotal = requestTotal
	rep.ValidTotal = validTotal
	return rep, ret
}

func (l *LogApp) QPSPeak(uid uint64) (vo.QPSPeak, int) {
	var res vo.QPSPeak
	qpsPeak, err := l.domain.QPSPeak(uid)
	if err != nil {
		log.Logger.Error("get qps peak err", zap.Error(err))
		return res, interfaces.ErrorQPSPeak
	}
	res.QPSPeak = qpsPeak
	return res, interfaces.StatusSuccess
}

// NewLogApp new user app
func NewLogApp(repo *persistence.Repositories) LogApp {
	return LogApp{
		domain: domain.NewLogDomain(repo.APILog),
	}
}
