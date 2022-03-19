package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	// "unsafe"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 类似装饰器
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValue()
	t.Log(a)
	tt := timeSpent(slowFunc)
	v := tt(10)
	t.Log(v)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestFnParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(2, 4, 3))
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) toString(t *testing.T) string {
	fmt.Printf("Address is %x\n", &e.Name)
	s := fmt.Sprintf("Id: %s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
	return s
}

func TestStruct(t *testing.T) {
	e1 := Employee{Name: "Mike", Age: 30}
	e := Employee{"0", "Bob", 20}

	fmt.Printf("1Address is %x\n", &e.Name)

	t.Log(e1)
	t.Logf("e1 is %T", e1)

	e2 := new(Employee) // 返回指针

	e2.Name = "Rose"
	e2.Age = 22
	e2.Id = "222"

	t.Log(e)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
	t.Logf("77777777777777777778: %s", e.toString(t))

}
