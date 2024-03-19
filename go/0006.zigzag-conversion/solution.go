// Created by Ming Lee Ng at 2024/03/19 16:14
// leetgo: dev
// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([]string, numRows)
	direction := 1
	row := 0

	for _, char := range s {
		rows[row] += string(char)
		row += direction

		if row == 0 || row == numRows-1 {
			direction *= -1
		}
	}

	return strings.Join(rows, "")
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	numRows := Deserialize[int](ReadLine(stdin))
	ans := convert(s, numRows)

	fmt.Println("\noutput:", Serialize(ans))
}
