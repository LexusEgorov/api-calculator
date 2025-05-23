package calculator

import "testing"

func TestSumNums(t *testing.T) {
	nums := []float64{1.2, 5.3, 2.2}
	expected := 8.7
	res := sumNums(nums...)

	if res != expected {
		t.Errorf("not equal! got: %f expected %f", res, expected)
	}
}

func TestMultNums(t *testing.T) {
	nums := []float64{2, 5, 4}
	var expected float64 = 40
	res := multNums(nums...)

	if res != expected {
		t.Errorf("not equal! got: %f expected %f", res, expected)
	}
}
