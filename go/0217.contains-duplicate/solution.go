// Created by Ming Lee Ng at 2024/04/06 15:17
// leetgo: dev
// https://leetcode.com/problems/contains-duplicate/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)

	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = 1
	}

	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := containsDuplicate(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
