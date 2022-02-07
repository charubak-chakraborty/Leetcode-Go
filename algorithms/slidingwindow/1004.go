package slidingwindow

// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

// Example 1:

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
// Example 2:

// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Constraints:

// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.
// 0 <= k <= nums.length

//==================================Approach: Sliding Window==================================

// Complexity Analysis

// Time Complexity: O(N)O(N), where NN is the number of elements in the array. In worst case we might end up visiting every element of array twice, once by left pointer and once by right pointer.

// Space Complexity: O(1)O(1). We do not use any extra space.
//======================================================================================================
func LongestOnes(nums []int, k int) int {
	left := 0
	right := 0
	zeroes := 0
	longest := 0
	for right <= len(nums)-1 {
		if nums[right] == 0 {
			zeroes++
		}
		for zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}
			left++
		}
		longest = max(longest, right-left+1)
		right++
	}
	return longest
}
