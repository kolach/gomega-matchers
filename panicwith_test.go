package matchers_test

import (
	"errors"

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

	It("should return correct failure message", func() {
		m.Match(func() { panic("bar") })
		Ω(m.FailureMessage(nil)).To(Equal("Expect to panic with\n\tfoo\nbut panicked with\n\tbar"))
	})

	It("should return correct negate failure message", func() {
		m.Match(func() { panic("foo") })
		Ω(m.NegatedFailureMessage(nil)).To(Equal("Expect not to panic with\n\tfoo\nbut it did"))
	})

	It("should work", func() {
		Ω(func() { panic("foo") }).Should(PanicWith("foo"))
		Ω(func() { panic("bar") }).ShouldNot(PanicWith("foo"))
	})
})

var _ = Describe("PanicWithError", func() {
	var fooErr = errors.New("foo")

	It("should check with what panic happens", func() {
		m := PanicWithError(fooErr)
		Ω(m.Match(func() { panic(errors.New("foo")) })).To(BeTrue())
	})
})
