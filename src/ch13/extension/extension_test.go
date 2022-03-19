package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
	age int
}

func (p *Pet) Speak() {
	fmt.Print("...n")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(host)
}

type Dog struct {
	p *Pet // 复合
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("hello xulsup")
	// t.Log(dog.p.age)s
}

type Cat struct {
	Pet
}

func (c *Cat) Speak() {
	fmt.Print(" miao miao ")
}

func TestCat(t *testing.T) {
	cat := new(Cat)
	cat.age = 22
	// cat.Speak()
	cat.SpeakTo(" i am cat") // 不支持重载 内嵌的Pet还是用Pet的方法
	t.Log(cat.age)
}
