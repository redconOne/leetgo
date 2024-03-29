// Created by Ming Lee Ng at 2024/03/29 14:57
// leetgo: dev
// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countSubarrays(nums []int, k int) (ans int64) {
	var (
		largest       = 0
		left    int64 = 0
		count         = 0
		result  int64 = 0
	)

	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}

	for _, num := range nums {
		if num == largest {
			count++
		}

		for count >= k {
			if nums[left] == largest {
				count--
			}
			left++
		}

		result += left
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := countSubarrays(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
