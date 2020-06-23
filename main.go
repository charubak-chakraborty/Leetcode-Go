package main

import (
	"fmt"

	leetcodeArray "./leetcodeArray"
)

func main() {
	nums := []int{1, 2, 3, 4}
	output := leetcodeArray.RunningSum(nums)
	fmt.Print(output)
}
