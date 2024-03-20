// Created by Ming Lee Ng at 2024/03/20 13:25
// leetgo: dev
// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxArea(height []int) (ans int) {
	largest := 0
	for left, right := 0, len(height)-1; left < right; {
		width := right - left
		minHeight := height[left]
		if height[left] > height[right] {
			minHeight = height[right]
		}
		current := width * minHeight
		if current > largest {
			largest = current
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return largest
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := maxArea(height)

	fmt.Println("\noutput:", Serialize(ans))
}
