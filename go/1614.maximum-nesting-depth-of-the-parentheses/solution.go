// Created by Ming Lee Ng at 2024/04/03 19:55
// leetgo: dev
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxDepth(s string) (ans int) {
	count, max := 0, 0

	for _, c := range s {
		if c == '(' {
			count++
		}
		if c == ')' {
			count--
		}

		if count > max {
			max = count
		}
	}

	return max
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := maxDepth(s)

	fmt.Println("\noutput:", Serialize(ans))
}
