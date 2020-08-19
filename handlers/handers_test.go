package handlers

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("abc", testIsPalindromeFunc("abc", false))
	t.Run("1aaa1", testIsPalindromeFunc("aaa", true))
	t.Run("1aa a 1", testIsPalindromeFunc("aa a", true))
	t.Run("1a,a a 1", testIsPalindromeFunc("a,a a ", true))
	t.Run(" ", testIsPalindromeFunc(" ", true))
	t.Run("", testIsPalindromeFunc("", true))
}

func testIsPalindromeFunc(message string, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		actual := isPalindrome(message)
		if actual != expected {
			t.Errorf("Expected the boolean of %s to be %t but instead got %t!", message, expected, actual)
		}
	}
}
