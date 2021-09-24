package application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/infrastructure/persistence"
)

var (
	logApp application.LogApp
)

func init() {
	os.Setenv("env", "test")
	config.GetConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	log.InitLogger()
	logApp = application.NewLogApp(repo)
}

// TestLogApp_QPSPeak
func TestLogApp_QPSPeak(t *testing.T) {
	res, ret := logApp.QPSPeak(13)
	fmt.Println(res, ret)
}
