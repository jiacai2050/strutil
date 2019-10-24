package strutil_test

import (
	"strutil"
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello, world",
			expected: "dlrow ,olleh",
		},
		{
			input:    "hello, 世界",
			expected: "界世 ,olleh",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			if v := strutil.Reverse(tc.input); tc.expected != v {
				t.Errorf("expected %s, actual %s", tc.expected, v)
			}
		})
	}
}
