package leetcodeArray

// Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

// Example 1:

// Input: "Hello"
// Output: "hello"
// Example 2:

// Input: "here"
// Output: "here"
// Example 3:

// Input: "LOVELY"
// Output: "lovely"

//Revisit
func ToLowerCase(str string) string {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] <= 90 && r[i] >= 65 {
			r[i] += 32
		}
	}
	return string(r)
}
