package main

import (
	"ready-go/boot"
	"ready-go/pkg/logger"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 基础组件初始化
	if err := boot.InitApp(); err != nil {
		logger.Fatal(err)
	}

}
