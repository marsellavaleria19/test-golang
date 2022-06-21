package count_palindrom

import "testing"

func TestCountPalindromSuccess(t *testing.T) {
	result := CountPalindrom(1, 10)
	if result == "" {
		// error
		panic("Result must be number")
	}
}

func TestCountPalindromFailed(t *testing.T) {
	result := CountPalindrom(0, 10)
	if result == "" {
		// error
		panic("Result must be number")
	}
}
