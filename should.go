package should

import (
	"testing"

	"github.com/AlekSi/should/are"
)

func assert(tb testing.TB, f func() bool, msg string, args ...any) bool {
	tb.Helper()

	if f() {
		return true
	}

	if len(args) == 0 {
		tb.Error(msg)
	} else {
		tb.Errorf(msg, args...)
	}

	return false
}

func BeEqual[T any](tb testing.TB, actual T, expected any) bool {
	tb.Helper()

	f := func() bool { return are.Equal(actual, expected) }
	return assert(tb, f, "%v != %v", actual, expected)
}

func BeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) bool {
	tb.Helper()

	f := func() bool { return are.Equal(actual, expected) }
	return assert(tb, f, msg, args...)
}

func NotBeEqual[T any](tb testing.TB, actual T, expected any) bool {
	tb.Helper()

	f := func() bool { return !are.Equal(actual, expected) }
	return assert(tb, f, "%v == %v", actual, expected)
}

func NotBeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) bool {
	tb.Helper()

	f := func() bool { return !are.Equal(actual, expected) }
	return assert(tb, f, msg, args...)
}

func BeZero[T any](tb testing.TB, actual T) bool {
	tb.Helper()

	var expected T
	f := func() bool { return are.Equal(actual, expected) }
	return assert(tb, f, "%v != %v", actual, expected)
}

func BeZerof[T any](tb testing.TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	var expected T
	f := func() bool { return are.Equal(actual, expected) }
	return assert(tb, f, msg, args...)
}

func NotBeZero[T any](tb testing.TB, actual T) bool {
	tb.Helper()

	var expected T
	f := func() bool { return !are.Equal(actual, expected) }
	return assert(tb, f, "%v == %v", actual, expected)
}

func NotBeZerof[T any](tb testing.TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	var expected T
	f := func() bool { return !are.Equal(actual, expected) }
	return assert(tb, f, msg, args...)
}
