package webEngin

import (
	"errors"

	"1-competition-test-congling.zhou/app/config"
)

var _ WebEngin = (*repository)(nil)

// WebEngin 对外数据引擎
type WebEngin interface {
	GetAPIWeight(apiName string) (weight int, err error)
}

type repository struct {
	nowAPI map[string]int
}

// GetAPIWeight 获取API权重
func (r *repository) GetAPIWeight(apiName string) (weight int, err error) {
	weight, isOK := r.nowAPI[apiName]
	if !isOK {
		return 0, errors.New("此接口没有权重数据：" + apiName)
	}
	return weight, nil
}

// New 初始化 web引擎
func New(config *config.Config) (WebEngin, error) {
	// TODO: 待探测获取首次请求
	detectionRes := map[string]int{"/api/2/now": 1}
	return &repository{
		nowAPI: detectionRes,
	}, nil
}
