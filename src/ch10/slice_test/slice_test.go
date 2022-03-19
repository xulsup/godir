package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // 不定长数组
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(s1); i++ {
		s0 = append(s0, i)
		t.Log(len(s0), cap(s0))
	}
}
