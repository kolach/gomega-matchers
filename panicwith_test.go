package matchers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "github.com/kolach/gomega-matchers"
)

var _ = Describe("PanicWith", func() {
	var m types.GomegaMatcher

	BeforeEach(func() {
		m = PanicWith("foo")
	})

	It("should check with what panic happens", func() {
		Ω(m.Match(func() { panic("foo") })).To(BeTrue())
	})

	It("should work", func() {
		Ω(func() { panic("foo") }).Should(PanicWith("foo"))
		Ω(func() { panic("bar") }).ShouldNot(PanicWith("foo"))
	})
})
