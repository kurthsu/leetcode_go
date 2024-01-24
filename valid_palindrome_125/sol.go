package valid_palindrome_125

import (
	"regexp"
	"strings"
)

func validPalindrome(s string) bool {
	str := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	strLower := strings.ToLower(str)
	r := []rune(strLower)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	strReverse := string(r)
	return strLower == strReverse
}
