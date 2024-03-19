// Created by Ming Lee Ng at 2024/03/19 16:04
// leetgo: dev
// https://leetcode.com/problems/longest-palindromic-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestPalindrome(s string) string {
	for length := len(s) - 1; length >= 0; length-- {
		for start := 0; start+length < len(s); start++ {
			if isPalindrome(start, start+length, s) {
				return s[start : start+length+1]
			}
		}
	}

	return ""
}

func isPalindrome(l, r int, s string) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
