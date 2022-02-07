package slidingwindow

// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
//==================================Approach 1: Using Brute-Force==================================
// Time Complexity: O(N^2)

// Space Complexity: O(N)
//======================================================================================================
// func LengthOfLongestSubstring(s string) int {
// 	rs := []rune(s)
// 	if len(rs) <= 1 {
// 		return len(rs)
// 	}
// 	longest := 0
// 	for left := 0; left < len(rs); left++ {
// 		seenChars := make(map[rune]bool)
// 		currentLength := 0
// 		for right := left; right < len(rs); right++ {
// 			currentChar := rs[right]
// 			if !seenChars[currentChar] {
// 				currentLength++
// 				seenChars[currentChar] = true
// 				longest = max(currentLength, longest)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return longest
// }

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

//==================================Approach 1: Using Sliding Window==================================
// Time Complexity: O(N)

// Space Complexity: O(N)
//======================================================================================================

func LengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	if len(rs) <= 1 {
		return len(rs)
	}
	longest := 0
	seenChar := make(map[rune]int)
	left := 0
	for right := 0; right < len(rs); right++ {
		currentChar := rs[right]
		_, ok := seenChar[currentChar]
		if ok {
			prevSeenChar := seenChar[currentChar]
			if prevSeenChar >= left {
				left = prevSeenChar + 1
			}
		}

		seenChar[currentChar] = right
		longest = max(longest, right-left+1)
	}
	return longest
}
