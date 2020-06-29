package go_bench

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestConcurrencyAdd(t *testing.T){
	var count int32
	for i:=0;i<100;i++{
		go func() {
			count++
		}()
	}
	fmt.Println(count)
	count = 0
	for i:=0;i<100;i++{
		go func() {
			old := atomic.LoadInt32(&count)
			new := old+1
			for atomic.CompareAndSwapInt32(&count,old,new){
				break
			}
		}()
	}
	fmt.Println(count)
}