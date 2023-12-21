package main

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v3"
	"os"
)

type Configure struct {
	Mysql Mysql  `yaml:"mysql"`
	Port  string `yaml:"port"`
}

type Mysql struct {
	URL      string `yaml:"url"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Database string `yaml:"database"`
}

func initConfigure() Configure {
	dataBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", getWorkingDirPath(), "application.yaml"))
	if err != nil {

		return Configure{}
	}
	config := Configure{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		glog.Errorf("Parsing the yaml file failed: %s", err.Error())
		return Configure{}
	}

	return config
}

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
