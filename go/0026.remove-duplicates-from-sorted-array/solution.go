// Created by Ming Lee Ng at 2024/03/29 15:19
// leetgo: dev
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	left, right := 0, 0

	for ; right < len(nums); right++ {
		if right == len(nums)-1 || nums[right] != nums[right+1] {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
