package main

import (
	"fmt"
	"time"

	"acs/pkg"
	"acs/util"

	"github.com/kardianos/service"
	"github.com/lupc/go_service"
)

func main() {
	srvConfig := &service.Config{
		Name:        "ACS",
		DisplayName: "ACS自动清理服务",
		Description: "ACS自动清理服务",
	}

	_ = go_service.RunWithService(srvConfig, run)
}

func run() {
	util.GetLogger().Info("自动清理服务启动..")

	var cfgPath = "./config.yml"
	var cfg = pkg.GetConfig(cfgPath)

	if cfg != nil && len(cfg.Configs) > 0 {
		for _, c := range cfg.Configs {
			if c.IsEnable {
				c := c
				util.GoWithRecover(fmt.Sprintf("清理任务[%v]出错", c.Dir), func() {
					for {
						pkg.CleanProcess(c)
						time.Sleep(c.CheckInterval)
					}
				})
			}
		}
	}
}
