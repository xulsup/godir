package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	t.Log(m1[2])
	t.Logf("len m1: %d", len(m1))
	m2 := map[int]int{}
	t.Logf("len m2: %d", len(m2))
	m3 := make(map[int]int, 2)
	t.Logf("len m3: %d", len(m3)) // len取的不是capacity

}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	if _, ok := m1[3]; ok {
		t.Log("key three existing")
	} else {
		t.Log("key three not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[3] = true
	n := 4
	if mySet[n] {
		t.Logf("%d exists", n)
	} else {
		t.Logf("not exists: %d", n)
	}
	n = 3
	delete(mySet, n)
	if mySet[n] {
		t.Logf("%d exists", n)
	} else {
		t.Logf("not exists: %d", n)
	}
}
