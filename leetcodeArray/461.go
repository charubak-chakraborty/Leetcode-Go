package leetcodeArray

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

// Given two integers x and y, calculate the Hamming distance.

// Note:
// 0 ≤ x, y < 231.

// Example:

// Input: x = 1, y = 4

// Output: 2

// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑

// The above arrows point to positions where the corresponding bits are different.

//================================================Approach 2: Bit Shift================================================================
// Time Complexity: \mathcal{O}(1)O(1), since the Integer is of fixed size in Python and Java, the algorithm takes a constant time. For an Integer of 32 bit, the algorithm would take at most 32 iterations.

// Space Complexity: \mathcal{O}(1)O(1), a constant size of memory is used, regardless the input.

//================================================================================================================================================
// func hammingDistance(x int, y int) int {
// 	xor := x ^ y
// 	dist := 0
// 	for xor != 0 {
// 		if (xor & 1) > 0 {
// 			dist++
// 		}
// 		xor = xor >> 1
// 	}
// 	return dist
// }

//================================================Approach 3: Brian Kernighan's Algorithm================================================================
// Time Complexity: \mathcal{O}(1)O(1).

// Similar as the approach of bit shift, since the size (i.e. bit number) of integer number is fixed, we have a constant time complexity.

// However, this algorithm would require less iterations than the bit shift approach, as we have discussed in the intuition.

// Space Complexity: \mathcal{O}(1)O(1), a constant size of memory is used, regardless the input.
//================================================================================================================================================
func HammingDistance(x int, y int) int {
	xor := x ^ y
	dist := 0
	for xor != 0 {
		dist++
		xor = xor & (xor - 1)
	}
	return dist
}
