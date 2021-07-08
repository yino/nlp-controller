package database

import (
	"github.com/yino/nlp-controller/config"
)

type MysqlConfig struct {
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
	MysqlPort     string
	MysqlDb       string
}

func GetMysqlConf() *MysqlConfig {
	conf := config.GetConf()
	return &MysqlConfig{
		MysqlHost:     conf.MySql.Host,
		MysqlUser:     conf.MySql.User,
		MysqlPassword: conf.MySql.Password,
		MysqlPort:     conf.MySql.Port,
		MysqlDb:       conf.MySql.Db,
	}
}
