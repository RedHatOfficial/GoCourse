package factorial

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFactorial(t *testing.T) {
	Convey("0! should be equal 1", t, func() {
		So(Factorial(0), ShouldEqual, 1)
	})
}

func TestFactorial2(t *testing.T) {
	Convey("10! should be greater than 1", t, func() {
		So(Factorial(10), ShouldBeGreaterThan, 1)
	})
	Convey("10! should be between 1 and 10000000", t, func() {
		So(Factorial(10), ShouldBeBetween, 1, 10000000)
	})
}
