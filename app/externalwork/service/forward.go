package service

import (
	"fmt"

	"1-competition-test-congling.zhou/app/pkg/internet"
	"1-competition-test-congling.zhou/app/pkg/logger"
)

type forwardModel interface {
	ForwardRequests() (interface{}, error)
}

// ForwardRequests 请求转发业务
func (i *impl) ForwardRequests() (interface{}, error) {
	detectionInfo := i.GetAPIList()
	// TODO: api选择算法待接入
	url := fmt.Sprintf("%s%s", detectionInfo.Master, detectionInfo.APIs[0].Route)
	req := internet.HttpReq{
		Header:  map[string]string{"Authorization": "Basic conglin.zhou"},
		URL:     url,
		Method:  detectionInfo.APIs[0].Method,
		Timeout: 10,
	}
	var resp interface{}
	err := internet.SendHttpRequest(&req, &resp)
	if err != nil {
		logger.Error.Println("api:", url, "-->转发失败！", err)
		return nil, err
	}
	return resp, nil
}
