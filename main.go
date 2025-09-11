package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var opPriority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func isOperator(token string) bool {
	_, ok := opPriority[token]
	return ok
}

func toPostfix(tokens []string) ([]string, error) {
	var output []string
	var ops []string

	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			output = append(output, token)
		} else if isOperator(token) {
			for len(ops) > 0 && opPriority[ops[len(ops)-1]] >= opPriority[token] {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, token)
		} else {
			return nil, fmt.Errorf("invalid token: %s", token)
		}
	}
	for len(ops) > 0 {
		output = append(output, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return output, nil
}

func evalPostfix(postfix []string) (int, error) {
	var stack []int
	for _, token := range postfix {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				res = a / b
			}
			stack = append(stack, res)
		} else {
			return 0, fmt.Errorf("unknown operator: %s", token)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

func evaluateExpression(expr string) (int, error) {
	tokens := strings.Fields(expr)
	if len(tokens) == 0 {
		return 0, fmt.Errorf("empty expression")
	}
	postfix, err := toPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evalPostfix(postfix)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-calculator \"1 + 2 * 3\"")
		return
	}
	expr := os.Args[1]
	fmt.Println("Evaluating:", expr)
	result, err := evaluateExpression(expr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
