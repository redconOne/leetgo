// Created by Ming Lee Ng at 2024/03/19 15:43
// leetgo: dev
// https://leetcode.com/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	dict := make(map[int]int, 0)

	for idx, num := range nums {
		if pos, ok := dict[target-num]; ok {
			return []int{pos, idx}
		}
		dict[num] = idx
	}

	return []int{}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
