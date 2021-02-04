package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SystemConf *SystemConf `json:"system_config"`
}

type SystemConf struct {
	Port int `json:"port"`
}

// InitFromFile 从文件中读取配置信息
func InitFromFile(name string) (*Config, error) {
	bts, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(bts, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
