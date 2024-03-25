// Created by Ming Lee Ng at 2024/03/25 12:53
// leetgo: dev
// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findDuplicates(nums []int) (ans []int) {
	output := []int{}

	for _, num := range nums {
		idx := abs(num)
		if nums[idx-1] < 0 {
			output = append(output, idx)
		} else {
			nums[idx-1] *= -1
		}
	}

	return output
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
