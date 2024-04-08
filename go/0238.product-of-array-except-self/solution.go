// Created by Ming Lee Ng at 2024/04/08 14:12
// leetgo: dev
// https://leetcode.com/problems/product-of-array-except-self/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func productExceptSelf(nums []int) (ans []int) {
	leftProduct, rightProduct := 1, 1
	result := []int{}

	for _, num := range nums {
		result = append(result, leftProduct)
		leftProduct *= num
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := productExceptSelf(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
