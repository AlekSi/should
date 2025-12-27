package must

import (
	"testing"

	"github.com/AlekSi/should/are"
)

func assert(tb testing.TB, f func() bool, msg string, args ...any) {
	tb.Helper()

	if f() {
		return
	}

	if len(args) == 0 {
		tb.Fatal(msg)
	} else {
		tb.Fatalf(msg, args...)
	}
}

func BeEqual[T any](tb testing.TB, actual T, expected any) {
	tb.Helper()

	f := func() bool { return are.Equal(actual, expected) }
	assert(tb, f, "%v != %v", actual, expected)
}

func BeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) {
	tb.Helper()

	f := func() bool { return are.Equal(actual, expected) }
	assert(tb, f, msg, args...)
}

func NotBeEqual[T any](tb testing.TB, actual T, expected any) {
	tb.Helper()

	f := func() bool { return !are.Equal(actual, expected) }
	assert(tb, f, "%v == %v", actual, expected)
}

func NotBeEqualf[T any](tb testing.TB, actual T, expected any, msg string, args ...any) {
	tb.Helper()

	f := func() bool { return !are.Equal(actual, expected) }
	assert(tb, f, msg, args...)
}

func BeZero[T any](tb testing.TB, actual T) {
	tb.Helper()

	var expected T
	f := func() bool { return are.Equal(actual, expected) }
	assert(tb, f, "%v != %v", actual, expected)
}

func BeZerof[T any](tb testing.TB, actual T, msg string, args ...any) {
	tb.Helper()

	var expected T
	f := func() bool { return are.Equal(actual, expected) }
	assert(tb, f, msg, args...)
}

func NotBeZero[T any](tb testing.TB, actual T) {
	tb.Helper()

	var expected T
	f := func() bool { return !are.Equal(actual, expected) }
	assert(tb, f, "%v == %v", actual, expected)
}

func NotBeZerof[T any](tb testing.TB, actual T, msg string, args ...any) {
	tb.Helper()

	var expected T
	f := func() bool { return !are.Equal(actual, expected) }
	assert(tb, f, msg, args...)
}
