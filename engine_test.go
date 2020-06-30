package go_bench

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConcurrencyAdd(t *testing.T){
	countA := 0
	var wg sync.WaitGroup
	wg.Add(300)
	for i:=0;i<100;i++{
		go func(group *sync.WaitGroup) {
			countA++
			group.Done()
		}(&wg)
	}
	var countB int32
	for i:=0;i<100;i++{
		go func(group *sync.WaitGroup,count *int32) {
			for{
				value := *count
				if atomic.CompareAndSwapInt32(count,value,value+1){
					break
				}
			}
			group.Done()
		}(&wg,&countB)
	}

	var countC int32
	for i:=0;i<100;i++{
		go func(group *sync.WaitGroup,count *int32,i int) {
			atomic.AddInt32(count,1)
			group.Done()
		}(&wg,&countC,i)
	}
	wg.Wait()
	assert.Equal(t,false,100==countA)
	assert.Equal(t,true,100==countB)
	assert.Equal(t,true,100==countC)
}