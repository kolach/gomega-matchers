package matchers

import (
	"fmt"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// BeCausedBy checks if actual error is caused by expected one
func BeCausedBy(err error) types.GomegaMatcher { return &causedByMatcher{expected: err} }

type causedByMatcher struct{ expected error }

func (matcher *causedByMatcher) Match(actual interface{}) (bool, error) {
	err, ok := actual.(error)
	if !ok {
		return false, nil
	}
	return errors.Cause(err) == matcher.expected, nil
}

func (matcher *causedByMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#v\tto be caused by\n\t%#v", actual, matcher.expected)
}

func (matcher *causedByMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\tnot to be caused by\n\t%#v", actual, matcher.expected)
}
