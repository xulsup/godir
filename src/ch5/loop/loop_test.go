package loop

import "testing"

func TestWithLoop(t *testing.T) {
	n := 0
	for n < 5 {
		if n == 2 {
			t.Log("222")
		} else {
			t.Log(n)
		}
		n++
	}
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}
