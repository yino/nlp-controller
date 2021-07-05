package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"nlp/infrastructure"
)

type Config struct {
	MySql struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Db       string `yaml:"db"`
	}
}

func GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("./" + infrastructure.GetEnv())
	fmt.Println(yamlFile, err)
	if err != nil {
		fmt.Println(err.Error())
		panic("config 文件不存在")
	}
	var _config *Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		panic("获取config yaml失败" + err.Error())
	}

	return _config
}
