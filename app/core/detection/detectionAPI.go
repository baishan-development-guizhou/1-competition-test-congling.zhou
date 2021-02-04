package detection

import (
	"fmt"
	"time"

	"1-competition-test-congling.zhou/app/config"
	"1-competition-test-congling.zhou/app/externalwork/domain/webEngin"
	"1-competition-test-congling.zhou/app/pkg/internet"
	"1-competition-test-congling.zhou/app/pkg/logger"
)

type detectionResults struct {
	Host  int `json:"host"`
	Value int `json:"value"`
}

// Start 开启探测任务
func Start(w webEngin.WebEngin, conf *config.Config) {
	for {
		logger.Info.Println("探测请求已开启！")
		for i := 0; i < len(conf.Detecation.APIs); i++ {
			go func(api *config.APIDetails) {
				url := fmt.Sprintf("%s%s", conf.Detecation.Master, api.Route)
				req := internet.HttpReq{
					Header:  map[string]string{"Authorization": "Basic conglin.zhou"},
					URL:     url,
					Method:  api.Method,
					Timeout: conf.SystemConf.TimeOut,
				}
				respData := detectionResults{}
				err := internet.SendHttpRequest(&req, &respData)
				if err != nil {
					logger.Warning.Println("api:", req.URL, "-->请求失败！", err)
				}
				w.SetAPIWeight(api.Route, respData.Value)
			}(conf.Detecation.APIs[i])
		}
		time.Sleep(time.Second * time.Duration(conf.SystemConf.DetecationTime))
	}
}
