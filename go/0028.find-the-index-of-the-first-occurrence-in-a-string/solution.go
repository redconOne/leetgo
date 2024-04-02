// Created by Ming Lee Ng at 2024/04/02 17:01
// leetgo: dev
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strStr(haystack string, needle string) (ans int) {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	haystack := Deserialize[string](ReadLine(stdin))
	needle := Deserialize[string](ReadLine(stdin))
	ans := strStr(haystack, needle)

	fmt.Println("\noutput:", Serialize(ans))
}
