// Created by Ming Lee Ng at 2024/04/08 19:50
// leetgo: dev
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(numbers []int, target int) (ans []int) {
	for left, right := 0, len(numbers)-1; left < right; {
		current := numbers[left] + numbers[right]

		switch {
		case current < target:
			left++
		case current > target:
			right--
		default:
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numbers := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(numbers, target)

	fmt.Println("\noutput:", Serialize(ans))
}
