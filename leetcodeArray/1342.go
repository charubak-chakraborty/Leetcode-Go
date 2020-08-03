package leetcodeArray

// Given a non-negative integer num, return the number of steps to reduce it to zero. If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

// Example 1:

// Input: num = 14
// Output: 6
// Explanation:
// Step 1) 14 is even; divide by 2 and obtain 7.
// Step 2) 7 is odd; subtract 1 and obtain 6.
// Step 3) 6 is even; divide by 2 and obtain 3.
// Step 4) 3 is odd; subtract 1 and obtain 2.
// Step 5) 2 is even; divide by 2 and obtain 1.
// Step 6) 1 is odd; subtract 1 and obtain 0.
// Example 2:

// Input: num = 8
// Output: 4
// Explanation:
// Step 1) 8 is even; divide by 2 and obtain 4.
// Step 2) 4 is even; divide by 2 and obtain 2.
// Step 3) 2 is even; divide by 2 and obtain 1.
// Step 4) 1 is odd; subtract 1 and obtain 0.
// Example 3:

// Input: num = 123
// Output: 12

// Constraints:

// 0 <= num <= 10^6

//=====================================approach 1:Simulation=====================================
// Let n = num
// Time Complexity : O(\log \, n)O(logn).

// At each step, what we did depended on whether the remaining numnum was odd or even. If numnum was even, we halved what was left. If it was odd, we only subtracted 11. However, by subtracting 11, we were making it even, and so on the next step we were guaranteed to halve it.

// What this means is that in the worst case, we're halving it on every second step. We treat the \frac{1}{2}
// 2
// 1
// ​
//   of the time as a constant though, so in essence, we say that at each step, numnum is being halved.

// When something is halved at every step, it has a O(\log \, n)O(logn) time complexity.

// Space Complexity : O(1)O(1).

// We only use a constant number of integer variables, and so the space complexity is O(1)O(1).
// ===============================================================================================================
// func NumberOfSteps(num int) int {
// 	steps := 0
// 	for num != 0 {
// 		if num%2 == 0 {
// 			num = num / 2
// 		} else {
// 			num = num - 1
// 		}
// 		steps++
// 	}
// 	return steps
// }

//=====================================approach 2:bit count=====================================
// Let n = num

// Time Complexity : O(\log \, n)O(logn).

// Converting a number into string can be done in \log \, nlogn time.

// We then loop over each bit, doing a single operation each time. The number of bits in a number is \log_2 \, numberlog
// 2
// ​
//  number, so the time complexity is O(\log \, n)O(logn).

// Space Complexity : O(\log \, n)O(logn).

// Because we convert the number into a string, we'll have \log_2 \, numberlog
// 2
// ​
//  number characters in our string. This gives us a space complexity of O(\log \, n)O(logn).
//===============================================================================================================
// func NumberOfSteps(num int) int {
// 	steps := 0
// 	bin := strconv.FormatInt(int64(num), 2)
// 	for _, v := range bin {
// 		if string(v) == "1" {
// 			steps += 2
// 		} else {
// 			steps++
// 		}
// 	}
// 	return steps - 1
// }

//=====================================approach 3: Counting Bits with Bitwise Operators=====================================
// Let n = numn=num.

// Time Complexity : O(\log \, n)O(logn).

// We're pulling out each of the \log \, nlogn bits from num and performing an O(1)O(1) operation on each one. Therefore, the total time complexity is, again, O(\log \, n)O(logn).

// Space Complexity : O(1)O(1).

// We only use a constant number of integer variables, and so the space complexity is O(1)O(1).
//===============================================================================================================
func NumberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	steps := 0
	poweroftwo := 1
	for poweroftwo <= num {
		if (poweroftwo & num) != 0 {
			steps += 2
		} else {
			steps++
		}
		poweroftwo = poweroftwo * 2
	}
	return steps - 1
}
