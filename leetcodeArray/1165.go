package leetcodeArray

// There is a special keyboard with all keys in a single row.

// Given a string keyboard of length 26 indicating the layout of the keyboard (indexed from 0 to 25), initially your finger is at index 0. To type a character, you have to move your finger to the index of the desired character. The time taken to move your finger from index i to index j is |i - j|.

// You want to type a string word. Write a function to calculate how much time it takes to type it with one finger.

// Example 1:

// Input: keyboard = "abcdefghijklmnopqrstuvwxyz", word = "cba"
// Output: 4
// Explanation: The index moves from 0 to 2 to write 'c' then to 1 to write 'b' then to 0 again to write 'a'.
// Total time = 2 + 1 + 1 = 4.
// Example 2:

// Input: keyboard = "pqrstuvwxyzabcdefghijklmno", word = "leetcode"
// Output: 73

// Constraints:

// keyboard.length == 26
// keyboard contains each English lowercase letter exactly once in some order.
// 1 <= word.length <= 10^4
// word[i] is an English lowercase letter.
func CalculateTime(keyboard string, word string) int {
	keyMap := make(map[string]int)
	var steps int
	for i, v := range keyboard {
		keyMap[string(v)] = i
	}

	for i := 0; i < len(word)-1; i++ {
		if i == 0 {
			steps += Abs(keyMap[string(word[i])] - 0)
		}
		steps += Abs(keyMap[string(word[i+1])] - keyMap[string(word[i])])
	}
	return steps
}
