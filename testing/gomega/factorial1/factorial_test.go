package main

import (
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := Factorial(0)
	if result != 1 {
		t.Errorf("Expected that 0! == 1, but got %d instead", result)
	}
}

func TestFactorialForOne(t *testing.T) {
	result := Factorial(1)
	if result != 1 {
		t.Errorf("Expected that 1! == 1, but got %d instead", result)
	}
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := Factorial(5)
	if result <= 10 || result >= 10000 {
		t.Errorf("Expected that 5! == is between 10..10000")
	}
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	result := Factorial(20)
	if result <= 10 || result >= 10000 {
		t.Errorf("Expected that 20! == is between 10..10000")
	}
}

func TestFactorialForTen(t *testing.T) {
	result := Factorial(10)
	expected := int64(3628800)
	if result != expected {
		t.Errorf("Expected that 10! == %d, but got %d instead", expected, result)
	}
}

func TestFactorialForBigNumber(t *testing.T) {
	result := Factorial(20)
	if result <= 0 {
		t.Errorf("Expected that 20! > 0, but got negative number %d instead", result)
	}
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	result := Factorial(30)
	if result <= 0 {
		t.Errorf("Expected that 30! > 0, but got negative number %d instead", result)
	}
}
