// Created by Ming Lee Ng at 2024/04/08 13:15
// leetgo: dev
// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countStudents(students []int, sandwiches []int) (ans int) {
	zeroes, ones, remaining := 0, 0, len(students)

	for _, pref := range students {
		if pref == 0 {
			zeroes++
		} else {
			ones++
		}
	}

	for _, sandwich := range sandwiches {
		if sandwich == 0 {
			if zeroes == 0 {
				return remaining
			}
			zeroes--
		} else {
			if ones == 0 {
				return remaining
			}
			ones--
		}
		remaining--
	}

	return remaining
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	students := Deserialize[[]int](ReadLine(stdin))
	sandwiches := Deserialize[[]int](ReadLine(stdin))
	ans := countStudents(students, sandwiches)

	fmt.Println("\noutput:", Serialize(ans))
}
