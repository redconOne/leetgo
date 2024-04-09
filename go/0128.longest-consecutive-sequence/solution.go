// Created by Ming Lee Ng at 2024/04/08 19:08
// leetgo: dev
// https://leetcode.com/problems/longest-consecutive-sequence/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestConsecutive(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	longest := 1

	for i, count := 0, 0; i < len(nums); i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			count++
			if count > longest {
				longest = count
			}
		} else if nums[i] != nums[i-1] {
			count = 1
		}
	}

	return longest
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := longestConsecutive(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
