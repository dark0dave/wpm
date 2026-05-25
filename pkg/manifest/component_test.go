package manifest

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	testCases := []struct {
		input    *WeiduComponent
		expected string
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
		result := tc.input.ToLogString()
		if result != tc.expected {
			log.Fatalf("%s != %s", result, tc.expected)
		}
	}
}
