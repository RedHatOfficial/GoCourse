package accumulator

import (
	"github.com/DATA-DOG/godog"
)

func iHaveAnAccumulatorWith(arg1 int) error {
	return godog.ErrPending
}

func iAddToAccumulator(arg1 int) error {
	return godog.ErrPending
}

func theAccumulatedResultShouldBe(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an accumulator with (\d+)$`, iHaveAnAccumulatorWith)
	s.Step(`^I add (\d+) to accumulator$`, iAddToAccumulator)
	s.Step(`^the accumulated result should be (\d+)$`, theAccumulatedResultShouldBe)
}
