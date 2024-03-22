// Created by Ming Lee Ng at 2024/03/21 21:42
// leetgo: dev
// https://leetcode.com/problems/3sum-closest/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSumClosest(nums []int, target int) (ans int) {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	diff := math.Abs(float64(target - closest))

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			current := nums[i] + nums[left] + nums[right]
			currentDiff := math.Abs(float64(target - current))
			if current == target {
				return current
			}

			if currentDiff < diff {
				closest = current
				diff = currentDiff
			}

			if current > target {
				right--
			} else {
				left++
			}
		}
	}

	return closest
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := threeSumClosest(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
