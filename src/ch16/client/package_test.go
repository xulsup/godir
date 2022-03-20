package client

import (
	"ch16/series"
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(35))
	t.Log(series.Jsu(678))
}

func TestCm(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
