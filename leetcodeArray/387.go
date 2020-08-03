package leetcodeArray

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

// Examples:

// s = "leetcode"
// return 0.

// s = "loveleetcode"
// return 2.

// Note: You may assume the string contains only lowercase English letters.

func FirstUniqChar(s string) int {
	var sMap [26]int
	for _, v := range s {
		sMap[v-'a']++
	}
	for i, v := range s {
		if sMap[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
