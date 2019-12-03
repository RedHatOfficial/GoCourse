package factorial_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "factorial"
)

var _ = Describe("Factorial", func() {
	Context("For zero input", func() {
		It("should be one", func() {
			Expect(Factorial(0)).To(Equal(int64(1)))
		})
	})
})
