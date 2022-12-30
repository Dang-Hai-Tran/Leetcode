package richestcustomerwealth

import (
	"reflect"
	"testing"
)

func TestEmptyAccounts(t *testing.T) {
	accounts := [][]int{}
	res := maximumWealth(accounts)
	expt := 0
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("result = %v, expected = %v", res, expt)
	}
}
func TestOneAccounts(t *testing.T) {
	accounts := [][]int{{1, 2, 3}}
	res := maximumWealth(accounts)
	expt := 6
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("result = %v, expected = %v", res, expt)
	}
}
func TestTwoAccounts(t *testing.T) {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	res := maximumWealth(accounts)
	expt := 6
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("result = %v, expected = %v", res, expt)
	}
}

func TestThreeAccounts(t *testing.T) {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	res := maximumWealth(accounts)
	expt := 10
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("result = %v, expected = %v", res, expt)
	}
}
