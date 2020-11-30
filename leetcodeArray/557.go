package leetcodeArray

import "strings"

// Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:
// Input: "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Note: In the string, each word is separated by single space and there will not be any extra space in the string.
func ReverseWords(s string) string {
	tokens := strings.Split(s, " ")
	for idx, word := range tokens {

		bword := []byte(word)

		for i, j := 0, len(word)-1; i < len(word)/2; i, j = i+1, j-1 {
			bword[i], bword[j] = word[j], word[i]
		}
		tokens[idx] = string(bword)

	}
	s = strings.Join(tokens, " ")
	return s
}
