// Created by Ming Lee Ng at 2024/04/02 16:39
// leetgo: dev
// https://leetcode.com/problems/remove-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeElement(nums []int, val int) (ans int) {
	count := 0

	for _, num := range nums {
		if num == val {
			continue
		}
		nums[count] = num
		count++
	}

	return count
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := removeElement(nums, val)

	fmt.Println("\noutput:", Serialize(ans))
}
