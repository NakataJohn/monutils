# monutil server

监控功能的服务端，负责管理agent端，注册，获取监控信息；

## 配置

配置文件路径在代码的config.go中设置，默认是在当前main文件所在同级目录的etc目录中的server.yml

```yaml
server:
  listen: 0.0.0.0:8765
  hbtimeout: 1800

log_dir: "./logs"
```

server监听配置和日志配置；

## 集成

若需要集成server到已有的项目，需要将main.go中的main方法改为其他方法名，在项目的启动位置已协程的方式引用该方法。