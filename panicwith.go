package matchers

import (
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// PanicWith returns a matcher to check panic happened with certain expectations
func PanicWith(expect interface{}, cmpOpts ...cmp.Option) types.GomegaMatcher {
	return &panicWithMatcher{expect: expect, cmpOpts: cmpOpts}
}

// PanicWithError expects panic happens with error object
func PanicWithError(err error) types.GomegaMatcher {
	return PanicWith(err, cmp.Comparer(func(x, y error) bool {
		return x.Error() == y.Error()
	}))
}

// panicWithMatcher checks panic happens with certain expected data
type panicWithMatcher struct {
	actual  interface{}
	expect  interface{}
	cmpOpts []cmp.Option
}

// Match checks condition is met
func (matcher *panicWithMatcher) Match(f interface{}) (success bool, err error) {
	if f == nil {
		return false, fmt.Errorf("PanicWithMatcher expects a non-nil function argument")
	}

	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return false, fmt.Errorf("PanicWithMatcher expects a function.  Got:\n%s", format.Object(f, 1))
	}
	if !(fType.NumIn() == 0 && fType.NumOut() == 0) {
		return false, fmt.Errorf("PanicWithMatcher expects a function with no arguments and no return value.  Got:\n%s", format.Object(f, 1))
	}

	success = false
	defer func() {
		if e := recover(); e != nil {
			matcher.actual = e
			success = cmp.Equal(matcher.actual, matcher.expect, matcher.cmpOpts...)
		}
	}()

	reflect.ValueOf(f).Call([]reflect.Value{})

	return
}

// FailureMessage returns message on matcher failure
func (matcher *panicWithMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expect to panic with\n\t%+v\nbut panicked with\n\t%+v", matcher.expect, matcher.actual)
}

// NegatedFailureMessage returns mesage on negated case
func (matcher *panicWithMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expect not to panic with\n\t%+v\nbut it did", matcher.expect)
}
