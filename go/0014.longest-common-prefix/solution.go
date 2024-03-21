// Created by Ming Lee Ng at 2024/03/21 13:58
// leetgo: dev
// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, word := range strs {
		i := 0
		for ; i < len(word) && i < len(prefix) && prefix[i] == word[i]; i++ {
		}
		prefix = prefix[:i]
	}

	return prefix
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := longestCommonPrefix(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
