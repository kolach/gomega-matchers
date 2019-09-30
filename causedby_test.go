package matchers_test

import (
	. "github.com/kolach/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

var _ = Describe("CausedBy", func() {
	var (
		cause error
		m     types.GomegaMatcher
	)

	BeforeEach(func() {
		cause = errors.Errorf("Boom")
		m = BeCausedBy(cause)
	})

	Describe("Match", func() {
		It("should return true if error is a cause", func() {
			Ω(m.Match(cause)).Should(BeTrue())
			err1 := errors.Wrap(cause, "Error1")
			Ω(m.Match(err1)).Should(BeTrue())
			err2 := errors.Wrap(err1, "Error2")
			Ω(m.Match(err2)).Should(BeTrue())
		})

		It("should return false if error is not the cause", func() {
			cause := errors.Errorf("Ooops! That's me!")
			Ω(m.Match(cause)).Should(BeFalse())
		})

		It("should return false for non-error types", func() {
			Ω(m.Match("foo")).Should(BeFalse())
			Ω(m.Match(1)).Should(BeFalse())
			Ω(m.Match(nil)).Should(BeFalse())
			Ω(m.Match(true)).Should(BeFalse())
		})
	})

	Describe("FailureMessage", func() {
		It("should clarify expectations", func() {
			Ω(m.FailureMessage("Foo")).Should(Equal("Expected\n\t\"Foo\"\tto be caused by\n\tBoom"))
		})
	})

	Describe("FailureMessage", func() {
		It("should clarify expectations", func() {
			Ω(m.NegatedFailureMessage("Foo")).Should(Equal("Expected\n\t\"Foo\"\tnot to be caused by\n\tBoom"))
		})
	})

})
