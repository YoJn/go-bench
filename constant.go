package go_bench

import (
	"io"
	"os"
)

var (
	// defaultEngineCount 默认总次数
	defaultEngineCount int32 = 10000
	// defaultWorkerNum 默认协程数量
	defaultWorkerNum int32 = 100
	// defaultWriter 默认日志输出
    defaultWriter io.Writer = os.Stdout
)