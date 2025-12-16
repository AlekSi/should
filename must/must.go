package must

import (
	"testing"

	"github.com/AlekSi/should"
)

func BeEqual[T any](tb testing.TB, actual T, expected any) {
	tb.Helper()

	BeEqualf(tb, actual, expected, "%v != %v", actual, expected)
}

func BeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) {
	tb.Helper()

	if !should.BeEqualf(tb, actual, expected, msg, args...) {
		tb.FailNow()
	}
}

func BeZero[T any](tb testing.TB, actual T) {
	tb.Helper()

	var expected T
	BeEqual(tb, actual, expected)
}

func BeZerof[T any](tb testing.TB, actual T, msg string, args ...any) {
	tb.Helper()

	var expected T
	BeEqualf(tb, actual, expected, msg, args...)
}
