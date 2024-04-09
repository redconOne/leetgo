// Created by Ming Lee Ng at 2024/04/09 02:05
// leetgo: dev
// https://leetcode.com/problems/evaluate-reverse-polish-notation/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func evalRPN(tokens []string) (ans int) {
	if len(tokens) == 1 {
		result, err := strconv.Atoi(tokens[0])
		if err != nil {
			return -1
		}
		return result
	}

	for i := 2; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			result := parse(tokens, i)
			return evalRPN(result)
		}
	}
	return -2
}

func parse(tokens []string, idx int) []string {
	pre, post := tokens[0:idx-2], tokens[idx+1:]
	str := calculate(tokens[idx-2], tokens[idx-1], tokens[idx])
	remaining := append(pre, str)
	remaining = append(remaining, post...)
	return remaining
}

func calculate(str1, str2, operation string) string {
	num1, err1 := strconv.Atoi(str1)
	if err1 != nil {
		return "-1"
	}
	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		return "-1"
	}
	var result int
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	return strconv.Itoa(result)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tokens := Deserialize[[]string](ReadLine(stdin))
	ans := evalRPN(tokens)

	fmt.Println("\noutput:", Serialize(ans))
}
