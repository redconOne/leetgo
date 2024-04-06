// Created by Ming Lee Ng at 2024/04/06 13:26
// leetgo: dev
// https://leetcode.com/problems/count-alternating-subarrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countAlternatingSubarrays(nums []int) (ans int64) {
	var (
		total int64 = 0
		i           = 0
	)

	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[j-1] {
			i = j
		}
		total += int64(j - i)
	}

	return total + int64(len(nums))
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := countAlternatingSubarrays(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
