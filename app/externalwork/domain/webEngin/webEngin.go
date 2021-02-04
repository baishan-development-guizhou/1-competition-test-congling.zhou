package webEngin

import (
	"errors"
	"sync"

	"1-competition-test-congling.zhou/app/config"
)

var _ WebEngin = (*repository)(nil)

// WebEngin 对外数据引擎
type WebEngin interface {
	GetAPIWeight(apiName string) (weight int, err error)
	SetAPIWeight(apiName string, weight int)
	GetAPIList() *config.Detecation
}

type repository struct {
	sync    sync.Mutex
	nowAPI  map[string]int
	apiList *config.Detecation
}

// GetAPIWeight 获取API权重
func (r *repository) GetAPIWeight(apiName string) (weight int, err error) {
	weight, isOK := r.nowAPI[apiName]
	if !isOK {
		return 0, errors.New("此接口没有权重数据：" + apiName)
	}
	return weight, nil
}

func (r *repository) SetAPIWeight(apiName string, weight int) {
	r.sync.Lock()
	defer r.sync.Unlock()
	r.nowAPI[apiName] = weight
}

func (r *repository) GetAPIList() *config.Detecation {
	return r.apiList
}

// New 初始化 web引擎
func New(config *config.Config) (WebEngin, error) {
	repository := &repository{
		nowAPI:  make(map[string]int, config.SystemConf.APINum),
		apiList: config.Detecation,
	}
	return repository, nil
}
