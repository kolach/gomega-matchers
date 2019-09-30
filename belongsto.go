package matchers

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// BelongTo checks if actual error is caused by expected one
func BelongTo(a ...interface{}) *BelongToMatcher { return &BelongToMatcher{a: a} }

type BelongToMatcher struct {
	a       []interface{}
	cmpOpts []cmp.Option
}

func BelongToWithOpts(opts ...cmp.Option) func(a ...interface{}) *BelongToMatcher {
	return func(a ...interface{}) *BelongToMatcher {
		return &BelongToMatcher{a, opts}
	}
}

func (matcher *BelongToMatcher) WithOpts(opts ...cmp.Option) *BelongToMatcher {
	return &BelongToMatcher{matcher.a, opts}
}

func (matcher *BelongToMatcher) Match(actual interface{}) (bool, error) {
	for _, expected := range matcher.a {
		if cmp.Equal(actual, expected, matcher.cmpOpts...) {
			return true, nil
		}
	}
	return false, nil
}

func (matcher *BelongToMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#v\tto belong to\n\t%#v", actual, matcher.a)
}

func (matcher *BelongToMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\tnot to belong to\n\t%#v", actual, matcher.a)
}
