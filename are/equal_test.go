package are

import "testing"

func TestEqual(t *testing.T) {
	for name, tc := range map[string]struct {
		f        func() bool
		expected bool
	}{
		"Int":      {f: func() bool { return Equal(42, 42) }, expected: true},
		"IntInt32": {f: func() bool { return Equal(42, int32(42)) }, expected: false},
		"Int32Int": {f: func() bool { return Equal(int32(42), 42) }, expected: false},
	} {
		t.Run(name, func(t *testing.T) {
			if a := tc.f(); a != tc.expected {
				t.Errorf("actual %v, expected %v", a, tc.expected)
			}
		})
	}
}
