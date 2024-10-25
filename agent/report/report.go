package report

import (
	"monutil/agent/config"
	"monutil/agent/utils"
)

func TimeTask(cfg *config.Conf, cli *Client) {
	utils.MonitorTask(cfg, cli)
}
