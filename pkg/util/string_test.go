package util

import (
	"testing"
)

// Tests the RandString that generates random hex strings.
//
// It provides check for the correct length and then checks that the
// characters are valid hex numbers (0-9A-F).
// For this implementation it does not consider hex strings with lower case
// characters (a-f).
func TestRandString(t *testing.T) {
	for _, i := range []int{0, 1, 10, 40} {
		hexString := RandString(i)
		strLength := len(hexString)
		if strLength != i {
			t.Errorf("RandString(%d) returned string of length %d", i, strLength)
		} else if strLength > 0 {
			for _, c := range hexString {
				if !IsValidHexChar(c) {
					t.Errorf("RandString(%d) returned invalid hex number: %s", i, hexString)
				}
			}
		}
	}
}

// Tests that the rune provided is a valid hex number (0-9A-F)
func IsValidHexChar(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	} else if c >= 'A' && c <= 'F' {
		return true
	}
	return false
}
