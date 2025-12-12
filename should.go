package should

import (
	"testing"

	"github.com/AlekSi/should/are"
)

func BeEqual[T any](tb testing.TB, actual T, expected any) bool {
	tb.Helper()

	return BeEqualf(tb, actual, expected, "%v != %v", actual, expected)
}

func BeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) bool {
	tb.Helper()

	if !are.Equal(actual, expected) {
		tb.Errorf(msg, args...)
		return false
	}

	return true
}
