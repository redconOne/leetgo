// Created by Ming Lee Ng at 2024/03/20 13:54
// leetgo: dev
// https://leetcode.com/problems/integer-to-roman/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func intToRoman(num int) string {
	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX",
		"V", "IV", "I",
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for idx, val := range values {
		for num >= val {
			result += symbols[idx]
			num -= val
		}
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[int](ReadLine(stdin))
	ans := intToRoman(num)

	fmt.Println("\noutput:", Serialize(ans))
}
