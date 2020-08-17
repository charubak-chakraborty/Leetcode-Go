package leetcodeArray

// Given two strings s and t , write a function to determine if t is an anagram of s.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false
// Note:
// You may assume the string contains only lowercase alphabets.

func IsAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
