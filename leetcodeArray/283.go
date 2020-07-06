package leetcodeArray

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Example:

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Note:

// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
func MoveZeroes(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			for j := i; j > 0; j-- {
				if nums[j-1] != 0 {
					break
				} else {
					temp := nums[j-1]
					nums[j-1] = nums[j]
					nums[j] = temp
				}
			}
		}
	}
	return nums
}

//Revisit
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
