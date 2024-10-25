# monutil agent

agent当部署在需要采集监控指标的服务器上，不同的系统类型可以编译为对应的二进制包。

## 配置

配置文件路径在代码的config.go中设置，默认是在当前main文件所在同级目录的etc目录中的agent.yml

```yaml
agent:
  name: "test"                  # agent的注册名，必须唯一，无法重复注册相同name的agent。

server:
  host: 192.168.145.199:8765    # 服务端监听地址

heartbeat:                      # 心跳设置
  retry: 5  
  timeout: 30

monitor:                        # 监控设置，间隔和包含排除的监控项
  # eg:10s;3m;1h
  interval: 30s
  # include默认是所有：[cpu,disk,mem,host,load,net]。
  include: ["host", "cpu", "mem", "disk", "load"]
  exclude: ["net", "cpu", "disk"]

```

服务器端采集的指标请参考：[gopsutil](https://pkg.go.dev/github.com/shirou/gopsutil/v3#section-readme)