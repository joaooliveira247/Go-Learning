package math

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 5 {
		t.Error("Expect 5, got", v)
	}
}

func TestDouble(t *testing.T) {
	v := Double(2)
	if v != 4 {
		t.Error("Expect 4, got", v)
	}
}