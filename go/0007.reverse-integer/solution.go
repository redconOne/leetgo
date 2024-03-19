// Created by Ming Lee Ng at 2024/03/19 16:18
// leetgo: dev
// https://leetcode.com/problems/reverse-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverse(x int) (ans int) {
	const (
		lowBound  = -(1 << 31)
		highBound = (1 << 31) - 1
	)
	result, sign, num := 0, 1, x
	if x < 0 {
		sign *= -1
		num *= -1
	}

	for num != 0 {
		result = result*10 + num%10
		num = num / 10
	}

	result *= sign

	if result < lowBound || result > highBound {
		return 0
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := reverse(x)

	fmt.Println("\noutput:", Serialize(ans))
}
