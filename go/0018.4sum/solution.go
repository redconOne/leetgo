// Created by Ming Lee Ng at 2024/03/27 14:43
// leetgo: dev
// https://leetcode.com/problems/4sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		// empty for that prevents re-use of numbers
		for ; i > 0 && i < len(nums) && nums[i] == nums[i-1]; i++ {
		}

		for j := i + 1; j < len(nums); j++ {
			// empty for that prevents re-use of numbers
			for ; j > i+1 && j < len(nums) && nums[j] == nums[j-1]; j++ {
			}

			start, end := j+1, len(nums)-1

			for start < end {
				sum := nums[start] + nums[end] + nums[i] + nums[j]
				if sum < target {
					start++
					for ; start < len(nums) && nums[start] == nums[start-1]; start++ {
					}
				} else if sum > target {
					end--
					for ; end > 0 && nums[end] == nums[end+1]; end-- {
					}
				} else {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for ; start < len(nums) && nums[start] == nums[start-1]; start++ {
					}
					for ; end > 0 && nums[end] == nums[end+1]; end-- {
					}
				}
			}
		}
	}

	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := fourSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
