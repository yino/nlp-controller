package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yino/nlp-controller/infrastructure"

	"gopkg.in/yaml.v2"
)

const (
	PROD = "prod"
	DEV  = "dev"
	TEST = "test"
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
		QaHost      string `yaml:"qa_host"`
	}
	Log struct {
		Path     string `yaml:"path"`
		LogLevel string `yaml:"logLevel"`
	}
}

var Conf *Config

func GetConf() *Config {

	yamlFile, err := ioutil.ReadFile(infrastructure.GetEnv())
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

func GetEnv() string {
	env := os.Getenv("ENV")
	envMap := make(map[string]string)
	envMap[PROD] = PROD
	envMap[DEV] = DEV
	envMap[TEST] = TEST

	if v, ok := envMap[env]; !ok {
		return DEV
	} else {
		return v
	}
}
