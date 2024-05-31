package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {

	input := "{[]}}"

	fmt.Print(isValid(input))
	fmt.Print("\n")
	fmt.Print(isValid(")[]"))
}
