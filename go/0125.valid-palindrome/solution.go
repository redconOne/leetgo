// Created by Ming Lee Ng at 2024/04/08 19:38
// leetgo: dev
// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(s string) bool {
	str := strings.Map(convert, s)

	for left, right := 0, len(str)-1; left < right; left, right = left+1, right-1 {
		if str[left] != str[right] {
			return false
		}
	}

	return true
}

func convert(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}

	return unicode.ToLower(r)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
