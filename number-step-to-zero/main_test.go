package numbersteptozero

import (
	"testing"
	"reflect"
)

func TestNumber0(t *testing.T) {
	res := numberOfSteps(0)
	expt := 0
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("Result = %v, want %v", res, expt)
	}
}

func TestNumber14(t *testing.T) {
	res := numberOfSteps(14)
	expt := 6
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("Result = %v, want %v", res, expt)
	}
}

func TestNumber8(t *testing.T) {
	res := numberOfSteps(8)
	expt := 4
	if reflect.DeepEqual(res, expt) == false {
		t.Errorf("Result = %v, want %v", res, expt)
	}
}