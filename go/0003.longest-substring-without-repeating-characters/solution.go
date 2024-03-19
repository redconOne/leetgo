// Created by Ming Lee Ng at 2024/03/19 15:56
// leetgo: dev
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	dict := make(map[byte]int, 0)
	longest := 0

	for left, right := 0, 0; right < len(s); {
		if _, ok := dict[s[right]]; !ok || left == right {
			dict[s[right]] = 1
			right++
		} else {
			delete(dict, s[left])
			left++
		}
		if longest < right-left {
			longest = right - left
		}
	}

	return longest
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
