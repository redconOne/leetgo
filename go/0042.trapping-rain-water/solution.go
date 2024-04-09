// Created by Ming Lee Ng at 2024/04/08 20:11
// leetgo: dev
// https://leetcode.com/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {
	water := 0
	leftPoint, rightPoint := 0, len(height)-1
	leftMax, rightMax := height[leftPoint], height[rightPoint]

	for leftPoint < rightPoint {
		if leftMax < rightMax {
			leftPoint++
			if height[leftPoint] > leftMax {
				leftMax = height[leftPoint]
			}
			water += leftMax - height[leftPoint]
		} else {
			rightPoint--
			if height[rightPoint] > rightMax {
				rightMax = height[rightPoint]
			}
			water += rightMax - height[rightPoint]
		}
	}

	return water
}

func calcArea(width, height, obstacles int) int {
	return width*height - obstacles
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
