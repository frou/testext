package testext

import (
	"testing"
)

// HookedSubtestRunner returns a function that is equivalent to the Run(...)
// method of t, except that, in addition, the specified functions will be
// called before and after each subtest run using it.
//
// Pass a nil function to omit either one.
func HookedSubtestRunner(t *testing.T, beforeEach func(), afterEach func()) func(string, func(*testing.T)) {
	t.Helper()
	if beforeEach == nil && afterEach == nil {
		t.Log("Unnecessary use of `HookedSubtestRunner`. Use `t.Run` directly instead")
	}

	return func(name string, subtest func(*testing.T)) {
		t.Helper()
		if beforeEach != nil {
			t.Log("Calling hook [beforeEach]")
			beforeEach()
		}
		t.Run(name, subtest)
		if afterEach != nil {
			t.Log("Calling hook [afterEach]")
			afterEach()
		}
	}
}
