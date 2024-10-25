# Monutil



使用go-netty框架和gopsutil实现的一个主机各项指标监控采集上报共能的server-agent。



## install

拉取代码，并在go环境中进行安装和编译；

```bash
go mod tidy
# 服务端
cd server
go build
# 采集端
cd agent
go build

```

或者可以将服务端集成至已有的项目中，使用协程启动服务监听即可。

服务器端采集的指标请参考：[gopsutil](https://pkg.go.dev/github.com/shirou/gopsutil/v3#section-readme)