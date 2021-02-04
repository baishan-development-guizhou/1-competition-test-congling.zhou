package controller

import "1-competition-test-congling.zhou/app/externalwork/service"

var _ Interface = (*controller)(nil)

type controller struct {
	service service.Interface
}

// Interface 控制器抽象接口
type Interface interface {
}

// New 注册控制器
func New(service service.Interface) Interface {
	return &controller{service: service}
}
