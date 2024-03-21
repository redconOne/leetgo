// Created by Ming Lee Ng at 2024/03/21 13:28
// leetgo: dev
// https://leetcode.com/problems/roman-to-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func romanToInt(s string) (ans int) {
	values := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}
	result := 0

	for idx, sym := range symbols {
		if len(s) == 0 {
			break
		}
		length := len(sym)
		for len(s) >= length && s[:length] == sym {
			result += values[idx]
			s = s[length:]
		}
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := romanToInt(s)

	fmt.Println("\noutput:", Serialize(ans))
}
