package channel_test

import (
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 1024)
	return "service"
}
func otherTask() {
	time.Sleep(time.Second * 5)
}

func AsyncService() chan string {
	ch := make(chan string)
	go func() {
		ret := service()
		ch <- ret
	}()
	return ch
}

func TestAsyncService(t *testing.T) {
	ch := AsyncService()
	otherTask()
	t.Log("执行完成!")
	t.Log(<-ch)
}
