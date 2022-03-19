package err_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be large than 2")

func GetFibonacci(n int) ([]uint64, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	list := []uint64{1, 1}

	for i := 2; i < n; i++ {
		list = append(list, list[i-1]+list[i-2])
	}
	return list, nil

}

var cerror = errors.New("my custom error")

func TestGetFibonacci(t *testing.T) {
	defer func() {
		t.Log("catching error")
		if err := recover(); err != nil {
			t.Error(err)
		} else {
			t.Log("no error")
		}
	}()
	if list, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("less than two")
		}
		t.Error(err)
	} else {
		t.Log(list)
	}
	// panic(cerror)
}
