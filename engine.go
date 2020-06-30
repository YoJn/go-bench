package go_bench

import (
	"context"
	"time"
)

/// BenchEngine
type BenchEngine struct {
	// 内部记录函数链位置
	index int
	// 函数链
	methods BenchChain
	// 设置单个过期时间
	workExpireTime *time.Duration
	// 设置Engine过期时间
	engineExpireTime *time.Duration
	// 总次数
	totalCount int32
	// 工人(协程)数量
	workerNum int32
	// 工作函数
	work func(ctx context.Context)
}

/// Use
func (engine *BenchEngine) Use(middleware ...BenchHandle){
	engine.methods =append(engine.methods,middleware...)
}

/// SetWorkExpireTime
func (engine *BenchEngine) SetWorkExpireTime(duration time.Duration) *BenchEngine{
	engine.workExpireTime = &duration
	return engine
}

/// SetWorkExpireTime
func (engine *BenchEngine) SetEngineExpireTime(duration time.Duration) *BenchEngine{
	engine.engineExpireTime = &duration
	return engine
}

/// Add
func (engine *BenchEngine) Add(work func(ctx context.Context)) *BenchEngine{
	engine.work = work
	return engine
}

/// Add
func (engine *BenchEngine) Run() *BenchEngine{
	if engine.work == nil{
		panic("engine worker can not be nil")
	}
}

/// New
func New(totalCount,workerNum int32) *BenchEngine{
	return &BenchEngine{
		index:      -1,
		methods:    BenchChain{},
		workExpireTime: nil,
		engineExpireTime: nil,
		totalCount:totalCount,
		workerNum:workerNum,
	}
}

/// Default
func Default() *BenchEngine{
	engine := New(defaultEngineCount,defaultWorkerNum)
	//todo log & recover
	engine.Use(DefaultLogger())
	return engine
}

type BenchHandle func(*BenchEngine)

type BenchChain []BenchHandle