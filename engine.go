package go_bench

import (
	"time"
)

/// BenchEngine
type BenchEngine struct {
	// 内部记录函数链位置
	index int
	// 函数链
	methods BenchChain
	// 设置过期时间
	expireTime *time.Duration
	// 总次数
	totalCount int32
}

/// Use
func (engine *BenchEngine) Use(middleware ...BenchHandle){
	engine.methods =append(engine.methods,middleware...)
}

/// SetExpireTime
func  (engine *BenchEngine) SetExpireTime(duration time.Duration) *BenchEngine{
	engine.expireTime = &duration
	return engine
}

/// New
func New() *BenchEngine{
	return &BenchEngine{
		index:      -1,
		methods:    BenchChain{},
		expireTime: nil,
	}
}

/// Default
func Default() *BenchEngine{
	engine := New()
	//todo log & recover
	engine.Use()
	return engine
}

type BenchHandle func(*BenchEngine)

type BenchChain []BenchHandle