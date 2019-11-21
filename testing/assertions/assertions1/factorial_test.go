package factorial_test

import (
	"factorial"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	result := factorial.Factorial(0)
	println(assertions.ShouldEqual(result, 1))
}

func TestFactorialForOne(t *testing.T) {
	result := factorial.Factorial(1)
	println(assertions.ShouldEqual(result, 1))
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := factorial.Factorial(5)
	println(assertions.ShouldBeBetween(result, 10, 10000))
}

func TestFactorialForSmallNegative(t *testing.T) {
	result := factorial.Factorial(20)
	println(assertions.ShouldBeBetween(result, 10, 10000))
}

func TestFactorialForTen(t *testing.T) {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	println(assertions.ShouldEqual(result, expected))
}

func TestFactorialForBigNumber(t *testing.T) {
	result := factorial.Factorial(20)
	println(assertions.ShouldBeGreaterThan(result, 0))
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	result := factorial.Factorial(30)
	println(assertions.ShouldBeGreaterThan(result, 0))
}
