// Created by Ming Lee Ng at 2024/03/27 15:21
// leetgo: dev
// https://leetcode.com/problems/valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValid(s string) bool {
	stack := []rune{}

	for _, sym := range s {
		if sym == '(' || sym == '{' || sym == '[' {
			stack = append(stack, sym)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !matchyMatchy(last, sym) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func matchyMatchy(str1, str2 rune) bool {
	switch str1 {
	case '(':
		return str2 == ')'
	case '[':
		return str2 == ']'
	case '{':
		return str2 == '}'
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
