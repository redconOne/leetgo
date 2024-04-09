// Created by Ming Lee Ng at 2024/04/09 03:18
// leetgo: dev
// https://leetcode.com/problems/daily-temperatures/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func dailyTemperatures(temperatures []int) (ans []int) {
	results := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}
		stack = append(stack, i)
	}

	return results
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	temperatures := Deserialize[[]int](ReadLine(stdin))
	ans := dailyTemperatures(temperatures)

	fmt.Println("\noutput:", Serialize(ans))
}
