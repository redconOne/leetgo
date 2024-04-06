// Created by Ming Lee Ng at 2024/04/06 12:56
// leetgo: dev
// https://leetcode.com/problems/harshad-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumOfTheDigitsOfHarshadNumber(x int) (ans int) {
	original, sum := x, 0

	for ; x > 0; sum, x = sum+x%10, x/10 {
	}

	if original%sum == 0 {
		return sum
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := sumOfTheDigitsOfHarshadNumber(x)

	fmt.Println("\noutput:", Serialize(ans))
}
