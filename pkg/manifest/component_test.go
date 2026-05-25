package manifest

import (
	"fmt"
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	testCases := []struct {
		input          *WeiduComponent
		name, expected string
	}{
		{
			input: &WeiduComponent{
				tpFile: "test.tp2",
				name:   "test",
			},
			expected: fmt.Sprintf("~test%ctest.tp2~ #0 #0 // : ", os.PathSeparator),
		},
		{
			input:    &WeiduComponent{},
			expected: fmt.Sprintf("~%c~ #0 #0 // : ", os.PathSeparator),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.ToLogString()
			if result != tc.expected {
				t.Fatalf("%s != %s", result, tc.expected)
			}
		})
	}
}
