package boot

import (
	"flag"
	"fmt"
	"ready-go/config"
	"ready-go/pkg/logger"
)

func InitApp() error {
	// 加载配置文件
	configPath := flag.String("c", "config/config.toml", "config file path")
	if err := config.InitConfig(*configPath); err != nil {
		return fmt.Errorf("加载配置文件出错, %v", err)
	}

	cfg := config.GetConfig()
	// 初始化日志包
	logger.InitLog(cfg.Log.Style, cfg.Log.Level)

	return nil
}
