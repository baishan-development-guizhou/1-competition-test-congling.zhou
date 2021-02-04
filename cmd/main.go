package main

import (
	"flag"
	"fmt"

	"1-competition-test-congling.zhou/app/config"
	"1-competition-test-congling.zhou/app/core/detection"
	"1-competition-test-congling.zhou/app/externalwork/controller"
	"1-competition-test-congling.zhou/app/externalwork/domain/webEngin"
	"1-competition-test-congling.zhou/app/externalwork/middleware"
	"1-competition-test-congling.zhou/app/externalwork/route"
	"1-competition-test-congling.zhou/app/externalwork/service"
	"1-competition-test-congling.zhou/app/pkg/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("config", "app/config/config.json", "config")

func main() {
	//初始化配置文件
	flag.Parse()
	conf, err := config.InitFromFile(*configFile)
	if err != nil {
		logger.Error.Println("配置文件加载异常：", err)
	}

	engin, err := webEngin.New(conf)
	if err != nil {
		logger.Error.Println(fmt.Sprintf("fail to create webEngin:%v", err))
	}

	ginEngin := gin.Default()
	pprof.Register(ginEngin)

	// 初始化 service
	service := service.New(engin)
	// 初始化 controller
	controller := controller.New(service)
	// 初始化 middleware
	middleware := middleware.New(engin)

	// 开启探测任务
	go detection.Start(engin, conf)

	// 路由注册
	router := route.New(controller, middleware)
	router.Install(ginEngin)

	logger.Error.Println(fmt.Sprintf("server stop with err: %v", ginEngin.Run(fmt.Sprintf(":%d", conf.SystemConf.Port))))
}
