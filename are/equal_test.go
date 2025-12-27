package are

import (
	"errors"
	"io"
	"os"
	"testing"
)

func TestEqual(t *testing.T) {
	for name, tc := range map[string]struct {
		f        func() bool
		expected bool
	}{
		"Int":      {f: func() bool { return Equal(42, 42) }, expected: true},
		"IntInt32": {f: func() bool { return Equal(42, int32(42)) }, expected: false},
		"Int32Int": {f: func() bool { return Equal(int32(42), 42) }, expected: false},

		"ErrorNil1": {f: func() bool { return Equal(error(nil), nil) }, expected: true},
		"ErrorNil2": {f: func() bool { return Equal(error(nil), error(nil)) }, expected: true},
		"ErrorNil3": {f: func() bool { return Equal(error(nil), errors.New("")) }, expected: false},
		"ErrorNil4": {f: func() bool { return Equal(errors.New(""), nil) }, expected: false},

		"ErrorEOF1": {f: func() bool { return Equal(io.EOF, io.EOF) }, expected: true},
		"ErrorEOF2": {f: func() bool { return Equal(io.EOF, nil) }, expected: false},
		"ErrorEOF3": {f: func() bool { return Equal(error(nil), io.EOF) }, expected: false},

		"ErrorPathError1": {f: func() bool { return Equal(new(os.PathError), new(os.PathError)) }, expected: true},
		"ErrorPathError2": {f: func() bool { return Equal((*os.PathError)(nil), (*os.PathError)(nil)) }, expected: true},
		"ErrorPathError3": {f: func() bool { return Equal((*os.PathError)(nil), nil) }, expected: false},
		"ErrorPathError4": {f: func() bool { return Equal(error(nil), (*os.PathError)(nil)) }, expected: false},
	} {
		t.Run(name, func(t *testing.T) {
			if a := tc.f(); a != tc.expected {
				t.Errorf("actual %[1]v (%[1]T), expected %[2]v (%[2]T)", a, tc.expected)
			}
		})
	}
}
