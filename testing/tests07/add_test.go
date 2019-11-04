package main

import (
	"fmt"
	"math"
	"testing"
)

type AddTest struct {
	x        int32
	y        int32
	expected int32
}

func checkAdd(t *testing.T, testInputs []AddTest) {
	for _, i := range testInputs {
		result := add(i.x, i.y)
		if result != i.expected {
			msg := fmt.Sprintf("%d + %d should be %d, got %d instead",
				i.x, i.y, i.expected, result)
			t.Error(msg)
		}
	}
}

func TestAddBasicValues(t *testing.T) {
	var addTestInput = []AddTest{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{2, 1, 3},
	}
	checkAdd(t, addTestInput)
}

func TestAddNegativeValues(t *testing.T) {
	var addTestInput = []AddTest{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{2, 1, 3},
		{2, -2, 0},
	}
	checkAdd(t, addTestInput)
}

func TestAddMinValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MinInt32, 0, math.MinInt32},
		{math.MinInt32, 1, math.MinInt32 + 1},
	}
	checkAdd(t, addTestInput)
}

func TestAddMaxValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MaxInt32, 0, math.MaxInt32},
		{math.MaxInt32, 1, math.MinInt32},
		{math.MaxInt32, math.MinInt32, -1},
	}
	checkAdd(t, addTestInput)
}

func TestAddMinMaxValues(t *testing.T) {
	var addTestInput = []AddTest{
		{math.MinInt32, 0, math.MinInt32},
		{math.MinInt32, 1, math.MinInt32 + 1},
		{math.MaxInt32, 0, math.MaxInt32},
		{math.MaxInt32, 1, math.MinInt32},
		{math.MaxInt32, math.MinInt32, -1},
	}
	checkAdd(t, addTestInput)
}
