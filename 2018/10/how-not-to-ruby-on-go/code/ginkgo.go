package user_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("your model behavior", func() {
	BeforeEach(func() {
		setUpModels()
		setUpServices()
	})
	Context("Price Breakdown", func() {
	Context("Price Breakdown", func() {
		JustBeforeEach(func() {
			result := getResultFromMock()
		})
		It("should return amount without currency", func() {
			Ω(result).Should(BeEquivalentTo("200,50"))
		})
	})
})


// Ginkgo suite for a `tests` package
func TestAPI(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "API Suite")
}
