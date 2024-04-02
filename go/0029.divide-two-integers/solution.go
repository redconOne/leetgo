// Created by Ming Lee Ng at 2024/04/02 17:09
// leetgo: dev
// https://leetcode.com/problems/divide-two-integers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func divide(dividend int, divisor int) (ans int) {
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	negative := (dividend < 0) != (divisor < 0)
	dividend = abs(dividend)
	divisor = abs(divisor)

	result := 0

	for dividend >= divisor {
		tempDivisor := divisor
		multiple := 1

		for ; dividend >= (tempDivisor << 1); tempDivisor, multiple = tempDivisor<<1, multiple<<1 {
		}

		dividend -= tempDivisor
		result += multiple
	}

	if negative {
		return -result
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	dividend := Deserialize[int](ReadLine(stdin))
	divisor := Deserialize[int](ReadLine(stdin))
	ans := divide(dividend, divisor)

	fmt.Println("\noutput:", Serialize(ans))
}
