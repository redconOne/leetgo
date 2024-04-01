// Created by Ming Lee Ng at 2024/04/01 15:59
// leetgo: dev
// https://leetcode.com/problems/length-of-last-word/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLastWord(s string) (ans int) {
	i, count := len(s)-1, 0

	for ; s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i, count = i-1, count+1 {
	}

	return count
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLastWord(s)

	fmt.Println("\noutput:", Serialize(ans))
}
