package should

import "testing"

func TestEqual(t *testing.T) {
	BeEqualf(t, 42, 42, "not equal")
}
