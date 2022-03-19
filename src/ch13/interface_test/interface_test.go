package interface_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWord() Code
}

type GoProgrammer struct {
}
type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord() Code {
	return "fmt.Println(\"Hello xulsup with golang \")"
}

func (g *JavaProgrammer) WriteHelloWord() Code {
	return "fmt.Println(\"Hello xulsup with java\")"
}

// interface 指针调用
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v \n", p, p.WriteHelloWord())
}

func TestInterface(t *testing.T) {
	var gp *GoProgrammer
	var jp *JavaProgrammer
	gp = new(GoProgrammer)
	jp = new(JavaProgrammer)
	writeFirstProgram(gp)
	writeFirstProgram(jp)
}
