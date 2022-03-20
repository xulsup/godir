package channel_test

import (
	"sync"
	"testing"
	"time"
)

// 使用锁对并发做控制
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	arr := make([]int, 0, 600)
	for i := 0; i < 500; i++ {
		go func() {
			mut.Lock()
			counter++
			arr = append(arr, counter)
			mut.Unlock()
			t.Logf("counter is %d \n", counter)
		}()
	}
	defer func() {
		t.Log(arr)
	}()
	time.Sleep(5 * time.Second)
}

// 并发控制之等待所有协程完成
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	var arr []int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mut.Lock()
			counter++
			arr = append(arr, counter)
			mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(arr)
}

// csp 并发控制。利用一个通道来完成并发控制
