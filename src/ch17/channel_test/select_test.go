package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Second * 1)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestTimeout(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Logf("result %s", ret)
	case <-time.After(time.Second * 2):
		t.Error("time out")
	}
}

func TestSelect(t *testing.T) {
	ret := <-AsyncService()
	t.Log(ret)
}

func TestCancelChannel(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		ch <- 2
		time.Sleep(time.Second * 1)
		close(ch)
	}()

	// go func() {
	for {
		if _, ok := <-ch; !ok {
			t.Log("canceled")
			break
		}
	}
	// }()
	// time.Sleep(time.Second * 2)
}

func TestingChDead(t *testing.T) {
	var ch1 chan string = make(chan string)
	func() {
		time.Sleep(time.Second * 2)
		ch1 <- "111"
	}()
	t.Log("c1 is", <-ch1)
}
