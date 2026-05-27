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
			name: "name, subComponent, componentName, tpfile test",
			input: &WeiduComponent{
				tpFile:        "test.tp2",
				name:          "test",
				subComponent:  "test",
				componentName: "test",
			},
			expected: fmt.Sprintf("~test%ctest.tp2~ #0 #0 // test -> test", os.PathSeparator),
		},
		{
			name: "name and tpfile test",
			input: &WeiduComponent{
				tpFile: "test.tp2",
				name:   "test",
			},
			expected: fmt.Sprintf("~test%ctest.tp2~ #0 #0", os.PathSeparator),
		},
		{
			name: "tpfile only test",
			input: &WeiduComponent{
				tpFile: "test.tp2",
			},
			expected: "~test.tp2~ #0 #0",
		},
		{
			name:     "empty test",
			input:    &WeiduComponent{},
			expected: "~~ #0 #0",
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
