package calculator_test

import (
	"testing"

	"github.com/art-es/example-clean-microservice/internal/domain/calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator.Doubling()", func() {
	doubling := new(calculator.Calculator).Doubling

	When("number is 5", func() {
		var num float64 = 5

		It("should return 10", func() {
			Expect(doubling(num)).Should(Equal(float64(10)))
		})
	})
})
