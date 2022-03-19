package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknown type")
}

func DoSomethingS(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("v", v)
	default:
		fmt.Println("unkown type")
	}
}

func TestEmptyInterfaceAssert(t *testing.T) {
	DoSomethingS(10)
	DoSomethingS("100")
}
