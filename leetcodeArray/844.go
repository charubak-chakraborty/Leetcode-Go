package leetcodeArray

// Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

// Note that after backspacing an empty text, the text will continue empty.

// Example 1:

// Input: S = "ab#c", T = "ad#c"
// Output: true
// Explanation: Both S and T become "ac".
// Example 2:

// Input: S = "ab##", T = "c#d#"
// Output: true
// Explanation: Both S and T become "".
// Example 3:

// Input: S = "a##c", T = "#a#c"
// Output: true
// Explanation: Both S and T become "c".
// Example 4:

// Input: S = "a#c", T = "b"
// Output: false
// Explanation: S becomes "c" while T becomes "b".
// Note:

// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S and T only contain lowercase letters and '#' characters

func BackspaceCompare(S string, T string) bool {
	bs := 0
	oS := ""
	oT := ""
	for i := len(S) - 1; i >= 0; i-- {
		if string(S[i]) == "#" {
			bs++
		} else {
			if bs == 0 {
				oS += string(S[i])
			} else {
				bs--
			}
		}
	}
	bs = 0
	for i := len(T) - 1; i >= 0; i-- {
		if string(T[i]) == "#" {
			bs++
		} else {
			if bs == 0 {
				oT += string(T[i])
			} else {
				bs--
			}
		}
	}
	return oS == oT
}
