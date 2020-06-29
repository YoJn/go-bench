package main

type BenchEngine struct {
	// 内部记录函数链位置
	index int
	// 函数链
	methods BenchChain

}

type BenchHandle func(*BenchEngine)

type BenchChain []BenchHandle

func main() {
	
}
