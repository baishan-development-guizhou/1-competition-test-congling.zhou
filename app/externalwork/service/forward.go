package service

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"1-competition-test-congling.zhou/app/pkg/internet"
	"1-competition-test-congling.zhou/app/pkg/logger"
)

type forwardModel interface {
	ForwardRequests() (interface{}, error)
}

// ForwardRequests 请求转发业务
func (i *impl) ForwardRequests() (interface{}, error) {
	detectionInfo := i.GetAPIList()
	result, _ := rand.Int(rand.Reader, big.NewInt(30))
	url := fmt.Sprintf("%s%s", detectionInfo.Master, detectionInfo.APIs[int(result.Int64())].Route)
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
