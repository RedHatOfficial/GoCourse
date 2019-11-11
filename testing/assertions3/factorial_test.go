package factorial_test

import (
	"factorial"
	. "github.com/mockupcode/gunit/assert"
	. "github.com/smartystreets/assertions"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(0)
	assert.Equal(ShouldEqual(result, 1), "")
}

func TestFactorialForOne(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(1)
	assert.Equal(ShouldEqual(result, 1), "")
}

func TestFactorialForSmallNumber(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(5)
	assert.Equal(ShouldBeBetween(result, 10, 10000), "")
}

func TestFactorialForSmallNegative(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(20)
	assert.Equal(ShouldBeBetween(result, 10, 10000), "")
}

func TestFactorialForTen(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(10)
	expected := int64(3628800)
	assert.Equal(ShouldEqual(result, expected), "")
}

func TestFactorialForBigNumber(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(20)
	assert.Equal(ShouldBeGreaterThan(result, 0), "")
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	assert := GetAssertion(t)
	result := factorial.Factorial(30)
	assert.Equal(ShouldBeGreaterThan(result, 0), "")
}
