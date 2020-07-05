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
	// nums := []int{8, 1, 2, 2, 3}
	// output := leetcodeArray.SmallerNumbersThanCurrent(nums)
	// fmt.Print(output)

	// 1313
	// nums := []int{1, 2, 3, 4}
	// output := leetcodeArray.DecompressRLElist(nums)
	// fmt.Print(output)

	// 1389
	// nums := []int{1, 2, 3, 4, 0}
	// index := []int{0, 1, 2, 3, 0}
	// output := leetcodeArray.CreateTargetArray(nums, index)
	// fmt.Print(output)

	// 1295
	// nums := []int{555, 901, 482, 1771}
	// output := leetcodeArray.FindNumbers(nums)
	// fmt.Print(output)

	// 1266
	// nums := [][]int{{3, 2}, {-2, -2}}
	// output := leetcodeArray.MinTimeToVisitAllPoints(nums)
	// fmt.Print(output)

	// 1502
	// nums := []int{3, 4, 1}
	// output := leetcodeArray.CanMakeArithmeticProgression(nums)
	// fmt.Print(output)

	// 561
	nums := []int{1, 4, 3, 2}
	output := leetcodeArray.ArrayPairSum(nums)
	fmt.Print(output)

}
