package utils

import (
	"monutil/agent/config"
	"monutil/agent/logger"
	"monutil/agent/timer"
	"strings"
)

// 指标采集定时任务
func MonitorTask(cfg *config.Conf, cli Report) {
	tm := timer.NewTimerTask()
	interval := cfg.Viper.GetString("monitor.interval")
	spec := "@every " + interval
	if strings.Contains(interval, "s") || strings.Contains(interval, "m") || strings.Contains(interval, "h") {
		logger.Infof("监控以开启，监测间隔为:%s", interval)

		tm.AddTaskByFunc("MonitorTask", spec, func() {
			go Monitors(cfg, cli)
		})
	} else {
		logger.Error("配置错误，监控间隔时间没有指定单位.")
		return
	}
}
