## 构建方式

### 本地运行

进入项目目录下运行下列命令

```sh
go run cmd/main.go
```

### 本地构建

打包应用

```sh
GOSUMDB=off CGO_ENABLED=0 GOOS=linux go build -o zcl cmd/main.go
```

运行应用

```sh
./zcl -config "你的配置文件路径"
```

## 配置

### 当前配置项

```json
{
    "system_config":{
        "port": 8099,
        "api_num": 30,
        "detecation_time": 100,
        "time_out": 5
    },
    "detecation":{
        "apis":[
            {
                "route":"/api/1/now",
                "method": "GET"
            }
        ],
        "master": "http://114.117.0.25"
    }
    
}
```

**配置解读**

+ system_config 
  + port ：gin 运行端口
  + detecation_time ：接口探测时间间隔 单位 秒
  + time_out ：请求超时时间（当前只是探测超时时间）单位 秒
+ detecation
  + apis ：需要探测的api 列表
  + master：api 目标地址

## 日志

### 日志配置

**配置地址**

现目前只支持代码配置：`app\pkg\logger\logger.go`

### 自定义日志种类

```go
//自定义日志格式
Info = log.New(io.MultiWriter(file, os.Stderr), "【INFO】: ", log.Ldate|log.Ltime|log.Lshortfile)
Error = log.New(io.MultiWriter(file, os.Stderr), "【ERROR】: ", log.Ldate|log.Ltime|log.Lshortfile)
Warning = log.New(io.MultiWriter(file, os.Stderr), "【WARNING】: ", log.Ldate|log.Ltime|log.Lshortfile)
```

### 记录日志

```go
// 在需要使用日志的地方
logger.Error.Println(errormsg)
logger.Warning.Println(errormsg)
logger.Info.Println(errormsg)
```

## 项目结构

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

