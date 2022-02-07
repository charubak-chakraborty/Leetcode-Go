package slidingwindow

// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the max sliding window.

// Example 1:

// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]
// Example 3:

// Input: nums = [1,-1], k = 1
// Output: [1,-1]
// Example 4:

// Input: nums = [9,11], k = 2
// Output: [11]
// Example 5:

// Input: nums = [4,-2], k = 2
// Output: [4]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length

//==================================Approach 1: Use a Hammer (Bruteforce)==================================

// Intuition

// The straightforward solution is to iterate over all sliding windows and find a maximum for each window. There are N - k + 1 sliding windows and there are k elements in each window, that results in a quite bad time complexity O(Nk).

// As you can imagine, this straightforward solution would result in TLE (Time Limit Exceed) exception.
// Complexity Analysis

// Time complexity : O(Nk), where N is number of elements in the array.

// Space complexity : O(N−k+1) for an output array.
//======================================================================================================

// func MaxSlidingWindow(nums []int, k int) []int {
// 	n := len(nums)
// 	output := []int{}

// 	if n*k == 0 {
// 		output = append(output, 0)
// 		return output
// 	}
// 	for i := 0; i < n-k+1; i++ {
// 		longest := math.MinInt16
// 		for j := i; j < i+k; j++ {
// 			longest = max(longest, nums[j])
// 		}
// 		output = append(output, longest)
// 	}
// 	return output
// }

//==================================Approach 2: Deque==================================

// Complexity Analysis

// Time complexity : \mathcal{O}(N)O(N), since each element is processed exactly twice - it's index added and then removed from the deque.

// Space complexity : \mathcal{O}(N)O(N), since \mathcal{O}(N - k + 1)O(N−k+1) is used for an output array and \mathcal{O}(k)O(k) for a deque.
//======================================================================================================
func MaxSlidingWindow(nums []int, k int) []int {
	var q []int
	var result []int
	for i, n := range nums {
		// while q is not empty and current num is greater than rightmost q pop rightmost q element
		for len(q) != 0 && n >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		//push current element
		q = append(q, i)

		//remove left val from window
		if q[0] == i-k {
			q = q[1:]
		}
		//append result when u reach the window
		if i >= k-1 {
			result = append(result, nums[q[0]])
		}
	}
	return result

}

func maximum(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
