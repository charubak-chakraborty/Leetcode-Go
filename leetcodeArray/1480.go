package leetcodeArray

func RunningSum(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			nums[i] = nums[i] + (nums[i-1])
		}
	}
	return nums
}
