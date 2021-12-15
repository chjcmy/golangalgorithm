package Calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 3) != -2 {
		t.Errorf("Errors in case 1 : Expected 4, Actual %d", Add(1, 3))
	}
	if Add(1, 0) != 1 {
		t.Errorf("Errors in case 2 : Expected 2, Actual %d", Add(1, 0))
	}
	if Add(1, -1) != 2 {
		t.Errorf("Errors in case 3 : Expected 0, Actual %d", Add(1, -1))
	}
}
