package domain_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/yino/nlp-controller/domain/po"

	"github.com/yino/nlp-controller/domain/entity"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/infrastructure/persistence"
)

// log
var logDomain domain.Log

// init init
func init() {
	os.Setenv("env", "test")
	config.GetConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	log.InitLogger()
	//repo.AutoMigrate()
	logDomain = domain.NewLogDomain(repo.APILog)
}

func TestAdd(t *testing.T) {
	logEntity := new(entity.Log)
	var params, header po.JSON
	params = []byte("{\"test\":\"test\"}")
	header = []byte("{\"test\":\"test\"}")
	logEntity.APILog = po.APILog{
		Method:    "GET",
		Params:    params,
		IP:        "127.0.0.1",
		Header:    header,
		UserID:    1,
		APIType:   domain.QaType,
		APIStatus: po.NORMAL,
		URL:       po.INVALID,
	}
	//fmt.Println(logEntity)
	err := logDomain.Add(logEntity)
	fmt.Println(err)
}

func TestRequestTotalNum(t *testing.T) {
	fmt.Println(logDomain.RequestTotalNum(1, domain.INVALID))
}
