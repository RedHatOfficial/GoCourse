package factorial_test

import (
	"factorial"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	if result != 1 {
		t.Errorf("Expected 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	if result != expected {
		t.Errorf("Expected 0! == %d, but got %d instead", expected, result)
	}
}
