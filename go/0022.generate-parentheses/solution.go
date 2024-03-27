// Created by Ming Lee Ng at 2024/03/27 16:13
// leetgo: dev
// https://leetcode.com/problems/generate-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateParenthesis(n int) (ans []string) {
	result := []string{}

	dfs(n, 0, 0, "", &result)

	return result
}

func dfs(n, open, closed int, current string, arr *[]string) {
	if open == n && closed == n {
		*arr = append(*arr, current)
		return
	}

	if open < n {
		dfs(n, open+1, closed, current+"(", arr)
	}

	if open > closed {
		dfs(n, open, closed+1, current+")", arr)
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateParenthesis(n)

	fmt.Println("\noutput:", Serialize(ans))
}
