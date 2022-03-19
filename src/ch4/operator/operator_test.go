package operator_test

import (
	"testing"
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	t.Log("xxxxx=\n")
	t.Log(a == b)
}
