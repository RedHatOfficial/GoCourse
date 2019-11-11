package factorial_test

import (
	"factorial"
	"github.com/jacobsa/oglematchers"
	"github.com/jacobsa/ogletest"
	"testing"
)

func TestOgletest(t *testing.T) {
	ogletest.RunTests(t)
}

type FactorialTest struct{}

func init() {
	ogletest.RegisterTestSuite(&FactorialTest{})
}

func (t *FactorialTest) FactorialForZero() {
	result := factorial.Factorial(0)
	ogletest.ExpectThat(result, oglematchers.Equals(1))
}

func (t *FactorialTest) FactorialForOne() {
	result := factorial.Factorial(1)
	ogletest.ExpectThat(result, oglematchers.Equals(1))
}

func (t *FactorialTest) TestFactorialSmallNumber() {
	result := factorial.Factorial(5)
	ogletest.ExpectThat(result, oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000)))
}

func (t *FactorialTest) TestFactorialSmallNumberNegative() {
	result := factorial.Factorial(20)
	ogletest.ExpectThat(result, oglematchers.AllOf(
		oglematchers.GreaterThan(10),
		oglematchers.LessThan(10000)))
}

func (t *FactorialTest) TestFactorialForTen() {
	result := factorial.Factorial(10)
	expected := int64(3628800)
	ogletest.ExpectThat(result, oglematchers.Equals(expected))
}

func (t *FactorialTest) TestFactorialBigNumber() {
	result := factorial.Factorial(20)
	ogletest.ExpectThat(result, oglematchers.GreaterThan(0))
}

func (t *FactorialTest) TestFactorialEvenBiggerNumber() {
	result := factorial.Factorial(30)
	ogletest.ExpectThat(result, oglematchers.GreaterThan(0))
}
