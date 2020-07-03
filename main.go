package main

import (
	"fmt"

	leetcodeArray "./leetcodeArray"
)

func main() {
	// 1480
	// nums := []int{1, 2, 3, 4}
	// output := leetcodeArray.RunningSum(nums)
	// fmt.Print(output)

	// 1470
	// nums := []int{1, 1, 2, 2}
	// output := leetcodeArray.Shuffle(nums, 2)
	// fmt.Print(output)

	// 1431
	// candies := []int{2, 3, 5, 1, 3}
	// extraCandies := 3
	// output := leetcodeArray.KidsWithCandies(candies, extraCandies)
	// fmt.Print(output)

	// 1486
	// output := leetcodeArray.XorOperation(5, 0)
	// fmt.Print(output)

	// 1365
	nums := []int{8, 1, 2, 2, 3}
	output := leetcodeArray.SmallerNumbersThanCurrent(nums)
	fmt.Print(output)
}
