package test_string

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	s = "\xE4\xB8\xAC"
	t.Log(s)
	s = "中"
	t.Log("-------------------------------")
	t.Log(len(s))
	c := []rune(s)
	t.Logf("999: %d", len(c))
	t.Logf("unicode %x", c[0])
	t.Logf("utf8 %x", c)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)

	t.Log(strings.Join(parts, "+"))

}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str:" + s)
	i, err := strconv.Atoi("10")
	if err == nil {
		t.Log(10 + i)
	}
}
