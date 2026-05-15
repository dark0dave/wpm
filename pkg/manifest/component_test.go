package manifest

import (
	"log"
	"testing"
)

func expected(t *testing.T) {
	testCases := []struct {
		input    *WeiduComponent
		expected string
	}{
		{
			input:    nil,
			expected: "@123 ~test string~",
		},
		{
			input:    nil,
			expected: "123 ~test string~",
		},
		{
			input:    nil,
			expected: "@123 test string~",
		},
		{
			input:    nil,
			expected: "",
		},
	}
	for _, tc := range testCases {
		result := tc.input.ToLogString()
		if result != tc.expected {
			log.Fatalf("%s != %s", result, tc.expected)
		}
	}
}
