package leetcodeArray

// Given a string S, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and return the new string.

// Example 1:

// Input: "leetcodeisacommunityforcoders"
// Output: "ltcdscmmntyfrcdrs"
// Example 2:

// Input: "aeiou"
// Output: ""

// Note:

// S consists of lowercase English letters only.
// 1 <= S.length <= 1000

func RemoveVowels(S string) string {
	var output string
	for _, v := range S {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		}
		output += string(v)
	}
	return output
}
