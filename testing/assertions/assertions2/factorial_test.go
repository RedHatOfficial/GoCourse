package factorial_test

import (
	"factorial"
	"github.com/mockupcode/gunit/assert"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	assertion := assert.GetAssertion(t)
	result := factorial.Factorial(0)
	assertion.Equal(assertions.ShouldEqual(result, 1), "")
}

func TestFactorialForOne(t *testing.T) {
	assertion := assert.GetAssertion(t)
	result := factorial.Factorial(1)
	assertion.Equal(assertions.ShouldEqual(result, 1), "")
}

func TestFactorialForSmallNumber(t *testing.T) {
	assertion := assert.GetAssertion(t)
	result := factorial.Factorial(5)
	assertion.Equal(assertions.ShouldBeBetween(result, 10, 10000), "")
}

func TestFactorialForSmallNegative(t *testing.T) {
	assertion := assert.GetAssertion(t)
	result := factorial.Factorial(20)
	assertion.Equal(assertions.ShouldBeBetween(result, 10, 10000), "")
}

func TestFactorialForTen(t *testing.T) {
	assertion := assert.GetAssertion(t)
	result := factorial.Factorial(10)
	expected := int64(3628800)
	assertion.Equal(assertions.ShouldEqual(result, expected), "")
}

func TestFactorialForBigNumber(t *testing.T) {
	assert := assert.GetAssertion(t)
	result := factorial.Factorial(20)
	assert.Equal(assertions.ShouldBeGreaterThan(result, 0), "")
}

func TestFactorialForEvenBiggerNumber(t *testing.T) {
	assert := assert.GetAssertion(t)
	result := factorial.Factorial(30)
	assert.Equal(assertions.ShouldBeGreaterThan(result, 0), "")
}
