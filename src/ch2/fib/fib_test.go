package fib

import "testing"
import "fmt"

func TestFibList(t *testing.T) {

	var (
		a int = 1
		b     = 1
	)
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func ExchangeValue(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log("----\n")
	t.Log(a, b)
}
