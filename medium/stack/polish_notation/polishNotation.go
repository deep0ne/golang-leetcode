package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	numStack := make([]int, 0)
	for _, token := range tokens {
		if numToken, err := strconv.Atoi(token); err == nil {
			numStack = append(numStack, numToken)
		} else {
			second := numStack[len(numStack)-1]
			first := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-2]
			switch token {
			case "+":
				numStack = append(numStack, first+second)
			case "-":
				numStack = append(numStack, first-second)
			case "*":
				numStack = append(numStack, first*second)
			case "/":
				numStack = append(numStack, first/second)
			}
		}
	}
	return numStack[len(numStack)-1]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"} // must be 9
	fmt.Println(evalRPN(tokens))
}
