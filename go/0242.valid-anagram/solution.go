// Created by Ming Lee Ng at 2024/04/06 15:24
// leetgo: dev
// https://leetcode.com/problems/valid-anagram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAnagram(s string, t string) bool {
	dict := make(map[rune]int)

	for _, v := range s {
		fmt.Println(v)
		dict[v]++
	}

	for _, v := range t {
		fmt.Println(v)
		dict[v]--
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isAnagram(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
