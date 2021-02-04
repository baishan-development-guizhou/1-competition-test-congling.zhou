package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SystemConf *SystemConf `json:"system_config"`
	Detecation *Detecation `json:"detecation"`
}

type SystemConf struct {
	Port           int   `json:"port"`
	APINum         int   `json:"api_num"`
	DetecationTime int64 `json:"detecation_time"`
	TimeOut        int   `json:"time_out"`
}

type Detecation struct {
	APIs   []*APIDetails `json:"apis"`
	Master string        `json:"master"`
}

type APIDetails struct {
	Route  string `json:"route"`
	Method string `json:"method"`
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
