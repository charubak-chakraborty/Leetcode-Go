package leetcodeArray

import "strings"

// Given two binary strings, return their sum (also a binary string).

// The input strings are both non-empty and contains only characters 1 or 0.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"
// Example 2:

// Input: a = "1010", b = "1011"
// Output: "10101"

// Constraints:

// Each string consists only of '0' or '1' characters.
// 1 <= a.length, b.length <= 10^4
// Each string is either "0" or doesn't contain any leading zero.

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	zeroes := ""

	for i := 0; i < len(a)-len(b); i++ {
		zeroes += "0"
	}

	b = zeroes + b
	sum := ""
	aDigits := strings.Split(a, "")
	bDigits := strings.Split(b, "")
	c := 0
	for i := len(aDigits) - 1; i >= 0; i-- {
		if aDigits[i] == "1" && bDigits[i] == "1" {
			if c == 1 {
				sum = "1" + sum
			} else {
				sum = "0" + sum
			}
			c = 1
		} else if aDigits[i] == "0" && bDigits[i] == "0" {
			if c == 1 {
				sum = "1" + sum
			} else {
				sum = "0" + sum
			}
			c = 0
		} else {
			if c == 1 {
				sum = "0" + sum
				c = 1
			} else {
				sum = "1" + sum
			}
		}
	}
	if c == 1 {
		sum = "1" + sum
	}
	return sum
}
