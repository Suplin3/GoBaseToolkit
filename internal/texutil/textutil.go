package texutil

import (
	"GoBaseToolkit/internal/collections"
	"strings"
	"unicode"
)

func RunesCounter(s string) int {
	out := []rune(s)
	return len(out)
}

func ReveseString(s string) string {
	out := []rune(s)
	out = collections.Reverse(out)
	return string(out)
}

func IsPalindrome(s string) bool {
	return s == ReveseString(strings.ToLower(s))
}

func ToCamelCase(s string) string {
	runeS := []rune(s)
	out := make([]rune, 0, len(runeS))
	flag := false
	for _, r := range runeS {
		if flag {
			out = append(out, unicode.ToUpper(r))
			flag = false
		} else if string(r) != " " {
			out = append(out, r)
		}
		if string(r) == " " {
			flag = true
		}
	}
	return string(out)
}

func Truncate(s string, maxLen int) string {
	rs := []rune(s)
	if maxLen+1 > len(rs)-1 {
		return s
	} else {
		return string(rs[:maxLen]) + "..."
	}
}
