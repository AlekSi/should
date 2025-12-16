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

	if are.Equal(actual, expected) {
		return true
	}

	if len(args) == 0 {
		tb.Error(msg)
	} else {
		tb.Errorf(msg, args...)
	}

	return false
}

func BeZero[T any](tb testing.TB, actual T) bool {
	tb.Helper()

	var expected T
	return BeEqual(tb, actual, expected)
}

func BeZerof[T any](tb testing.TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	var expected T
	return BeEqualf(tb, actual, expected, msg, args...)
}
