// Created by Ming Lee Ng at 2024/03/27 13:24
// leetgo: dev
// https://leetcode.com/problems/subarray-product-less-than-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	left, product, result := 0, 1, 0
	if k <= 1 {
		return result
	}

	for right, num := range nums {
		product *= num

		for product >= k {
			product /= nums[left]
			left++
		}

		result += right - left + 1
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := numSubarrayProductLessThanK(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
