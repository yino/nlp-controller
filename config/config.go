package config

import (
	"fmt"
	"github.com/yino/nlp-controller/infrastructure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	MySql struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Db       string `yaml:"db"`
	}
	App struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		TokenExpire int64  `yaml:"token_expire"`
	}
}

var Conf *Config

func GetConf() *Config {
	pwdPath, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwdPath+"/../" + infrastructure.GetEnv())
	if err != nil {
		fmt.Println(err.Error())
		panic("config 文件不存在")
	}
	var _config *Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		panic("获取config yaml失败" + err.Error())
	}
	Conf = _config
	return _config
}
