package service

import "1-competition-test-congling.zhou/app/externalwork/domain/webEngin"

var _ Interface = (*impl)(nil)

type impl struct {
	webEngin.WebEngin
}

// Interface 抽象接口
type Interface interface {
	forwardModel
}

// New 初始化
func New(webEngin webEngin.WebEngin, opts ...Option) Interface {
	newService := &impl{WebEngin: webEngin}
	for _, opt := range opts {
		opt(newService)
	}
	return newService
}

// Option 其他的公共模块注入
type Option func(*impl)
