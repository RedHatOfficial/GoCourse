// Seriál "Programovací jazyk Go"
//
// Test práce s řezy (slices).

package slices

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceManipulation(t *testing.T) {
	Convey("Given an empty slice", t, func() {
		s := []int{}

		Convey("The slice should be empty initially", func() {
			So(s, ShouldBeEmpty)
		})

		Convey("When an item is added", func() {
			s = append(s, 1)

			Convey("The slice should not be empty", func() {
				So(s, ShouldNotBeEmpty)
			})
			Convey("The slice length should be one", func() {
				So(len(s), ShouldEqual, 1)
			})
			Convey("And length should NOT be zero, of course", func() {
				So(len(s), ShouldNotEqual, 0)
			})
			Convey("When another item is added", func() {
				s = append(s, 2)

				Convey("The slice length should be two", func() {
					So(len(s), ShouldEqual, 2)
				})
				Convey("And length should NOT be zero, of course", func() {
					So(len(s), ShouldNotEqual, 1)
				})
			})

			Convey("Now the slice length should be one again", func() {
				So(len(s), ShouldEqual, 1)
			})
		})
		Convey("And now the slice should be empty again", func() {
			So(s, ShouldBeEmpty)
		})
	})
}
