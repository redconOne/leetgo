// Created by Ming Lee Ng at 2024/03/19 21:57
// leetgo: dev
// https://leetcode.com/problems/palindrome-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(x int) bool {
	result, num := 0, x

	for num > 0 {
		result = result*10 + num%10
		num /= 10
	}

	return result == x
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := isPalindrome(x)

	fmt.Println("\noutput:", Serialize(ans))
}
