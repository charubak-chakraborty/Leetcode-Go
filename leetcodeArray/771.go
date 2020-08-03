package leetcodeArray

// You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

// The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Example 1:

// Input: J = "aA", S = "aAAbbbb"
// Output: 3
// Example 2:

// Input: J = "z", S = "ZZ"
// Output: 0
// Note:

// S and J will consist of letters and have length at most 50.
// The characters in J are distinct.
//========================================================Approach #2: Hash Set [Accepted]========================================================
// Time Complexity: O(J\text{.length} + S\text{.length}))O(J.length+S.length)). The O(J\text{.length})O(J.length) part comes from creating J. The O(S\text{.length})O(S.length) part comes from searching S.

// Space Complexity: O(J\text{.length})O(J.length).
//==============================================================================================================================================================================

func numJewelsInStones(J string, S string) int {
	m := map[string]bool{}
	count := 0
	for _, v := range J {
		m[string(v)] = true
	}
	for _, v := range S {
		if m[string(v)] {
			count++
		}
	}
	return count
}
