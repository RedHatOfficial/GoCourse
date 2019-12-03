package factorial_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFactorial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorial Suite")
}
