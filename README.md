```
1-competition-test-congling.zhou
├─ .github
│  └─ workflows
│     └─ go.yml
├─ .gitignore
├─ app：程序主体文件
│  ├─ config：配置中心
│  │  └─ config.go
│  ├─ core：程序负载均衡核心
│  ├─ externalwork：对外工作核心
│  │  ├─ controller
│  │  ├─ domain
│  │  ├─ matedata
│  │  ├─ middleware
│  │  └─ service
│  └─ pkg：公共包
│     ├─ internet
│     │  └─ sendReq.go
│     └─ logger
├─ cmd：程序脚本文件
│  └─ main.go
├─ LICENSE
└─ README.md
```