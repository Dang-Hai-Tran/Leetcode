package sum1darray

import (
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	res := runningSum([]int{})
	expected := []int{}
	if reflect.DeepEqual(res, expected) == false {
		t.Errorf("result = %v, expected = %v", res, expected)
	}
}

func Test1Element(t *testing.T) {
	res := runningSum([]int{1})
	expected := []int{1}
	if reflect.DeepEqual(res, expected) == false {
		t.Errorf("result = %v, expected = %v", res, expected)
	}
}

func Test2Element(t *testing.T) {
	res := runningSum([]int{1, 2})
	expected := []int{1, 3}
	if reflect.DeepEqual(res, expected) == false {
		t.Errorf("result = %v, expected = %v", res, expected)
	}
}

