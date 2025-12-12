package are

import "reflect"

func Equal[T any](actual T, expected any) bool {
	return reflect.DeepEqual(actual, expected)
}
