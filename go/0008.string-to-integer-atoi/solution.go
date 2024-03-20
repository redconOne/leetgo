// Created by Ming Lee Ng at 2024/03/19 16:26
// leetgo: dev
// https://leetcode.com/problems/string-to-integer-atoi/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myAtoi(s string) (ans int) {
	const (
		lowBound  = -(1 << 31)
		highBound = (1 << 31) - 1
	)
	sign, result, trimmed := 1, 0, strings.TrimSpace(s)

	if trimmed == "" {
		return result
	}

	switch trimmed[0] {
	case '-':
		sign = -1
		trimmed = trimmed[1:]
	case '+':
		trimmed = trimmed[1:]
	}

	for _, char := range trimmed {
		if !unicode.IsDigit(char) {
			break
		}

		num := int(char - '0')
		result = result*10 + num

		if result > highBound {
			if sign < 0 {
				return lowBound
			}
			return highBound
		}
	}

	return result * sign
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := myAtoi(s)

	fmt.Println("\noutput:", Serialize(ans))
}
