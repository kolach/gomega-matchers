package matchers_test

import (
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "github.com/kolach/gomega-matchers"
)

type Foo struct {
	s string
	B bool
	i int
}

// introduce custom belongTo to ignore unexported fields in Foo struct
var belongTo = BelongToWithOpts(cmpopts.IgnoreUnexported(Foo{}))

var _ = Describe("BelongTo", func() {
	var (
		opts = []interface{}{"abc", true, 4, Foo{"bar", true, 100}}
		m    types.GomegaMatcher
	)

	BeforeEach(func() {
		m = belongTo(opts...)
	})

	It("should return true if value belongs to array", func() {
		for _, opt := range opts {
			Ω(m.Match(opt)).Should(BeTrue())
		}
	})

	It("should return false if value does not belong to array", func() {
		Ω(m.Match(false)).Should(BeFalse())
		Ω(m.Match("abcd")).Should(BeFalse())
		Ω(m.Match(1)).Should(BeFalse())
	})

	It("should work", func() {
		Ω("bar").ShouldNot(belongTo(opts...))
		Ω("abc").Should(belongTo(opts...))
	})
})
