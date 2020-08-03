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

	// 1389 - TODO
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
	// nums := []int{1, 4, 3, 2}
	// output := leetcodeArray.ArrayPairSum(nums)
	// fmt.Print(output)

	// 561
	// nums := []int{2, 2, 1, 1, 1, 2, 2}
	// output := leetcodeArray.MajorityElement(nums)
	// fmt.Print(output)

	// 283
	// nums := []int{0, 1, 0, 3, 12}
	// output := leetcodeArray.MoveZeroes(nums)
	// fmt.Print(output)

	// 122
	// num := []int{1, 2, 3, 4, 5}
	// op := leetcodeArray.MaxProfit(num)
	// fmt.Print(op)

	//1221
	// s := "RLLLLRRRLR"
	// op := leetcodeArray.BalancedStringSplit(s)
	// fmt.Print(op)

	//455
	// g := []int{1, 2}
	// s := []int{1, 2, 3}
	// op := leetcodeArray.FindContentChildren(g, s)
	// fmt.Print(op)

	//1108
	// ip := "1.1.1.1"
	// op := leetcodeArray.DefangIPaddr(ip)
	// fmt.Print(op)

	//709
	// ip := "Hello"
	// op := leetcodeArray.ToLowerCase(ip)
	// fmt.Print(op)

	//709
	// nums := []int{3, 0, 1}
	// op := leetcodeArray.MissingNumber(nums)
	// fmt.Print(op)

	//1281
	// num := 24
	// op := leetcodeArray.SubtractProductAndSum(num)
	// fmt.Print(op)

	// 937 - TODO

	// logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}
	// op := leetcodeArray.ReorderLogFiles(logs)
	// fmt.Print(op)

	// 2

	// list1 := leetcodeArray.ListNode{2, &leetcodeArray.ListNode{4, &leetcodeArray.ListNode{3, nil}}}
	// list2 := leetcodeArray.ListNode{5, &leetcodeArray.ListNode{6, &leetcodeArray.ListNode{4, nil}}}
	// // list1 := leetcodeArray.ListNode{5, nil}
	// // list2 := leetcodeArray.ListNode{5, nil}
	// op := leetcodeArray.AddTwoNumbers(&list1, &list2)

	// for op != nil {
	// 	fmt.Print(op.Val)
	// 	op = op.Next
	// }

	// // 206

	// list1 := leetcodeArray.ListNode{2, &leetcodeArray.ListNode{4, &leetcodeArray.ListNode{3, nil}}}

	// op := leetcodeArray.ReverseList(&list1)

	// for op != nil {
	// 	fmt.Print(op.Val)
	// 	op = op.Next
	// }
	// 206

	// list1 := leetcodeArray.ListNode{2, &leetcodeArray.ListNode{4, &leetcodeArray.ListNode{3, nil}}}
	// list2 := leetcodeArray.ListNode{2, &leetcodeArray.ListNode{4, &leetcodeArray.ListNode{3, nil}}}

	// op := leetcodeArray.MergeTwoLists(&list1, &list2)

	// for op != nil {
	// 	fmt.Print(op.Val)
	// 	op = op.Next
	// }

	//67
	// res := leetcodeArray.AddBinary("1010", "1011")
	// fmt.Print(res)

	//9
	// res := leetcodeArray.IsPalindrome(1001)
	// fmt.Print(res)

	//387
	// res := leetcodeArray.FirstUniqChar("loveleetcode")
	// fmt.Print(res)

	//1119
	res := leetcodeArray.RemoveVowels("loveleetcode")
	fmt.Print(res)

	//1165 - TODO
	// res := leetcodeArray.CalculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")
	// fmt.Print(res)

	//217 -- TODO
	// nums := []int{3, 0, 1, 1}
	// res := leetcodeArray.ContainsDuplicate(nums)
	// fmt.Print(res)

	//1342

	// res := leetcodeArray.NumberOfSteps(14)
	// fmt.Print(res)
}
