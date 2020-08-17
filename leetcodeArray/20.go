package leetcodeArray

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

// Example 1:

// Input: "()"
// Output: true
// Example 2:

// Input: "()[]{}"
// Output: true
// Example 3:

// Input: "(]"
// Output: false
// Example 4:

// Input: "([)]"
// Output: false
// Example 5:

// Input: "{[]}"
// Output: true
func IsValid(s string) bool {
	stack := []string{}
	m := map[string]string{}
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["
	for _, r := range s {
		switch string(r) {
		case "(", "{", "[":
			stack = append(stack, string(r))
		case ")", "}", "]":
			if len(stack) == 0 || (stack[len(stack)-1] != m[string(r)]) {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
