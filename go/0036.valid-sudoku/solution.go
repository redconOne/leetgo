// Created by Ming Lee Ng at 2024/04/08 14:31
// leetgo: dev
// https://leetcode.com/problems/valid-sudoku/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValidSudoku(board [][]byte) bool {
	squares := [][]byte{}
	for i := 0; i < len(board); i++ {
		col := []byte{}
		squares = append(squares, []byte{}, []byte{}, []byte{})

		for j := 0; j < len(board); j++ {
			col = append(col, board[j][i])
			squareIdx := (i/3)*3 + (j / 3)
			squares[squareIdx] = append(squares[squareIdx], board[j][i])
		}

		if !isValid(board[i]) {
			return false
		}
		if !isValid(col) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if !isValid(squares[i]) {
			return false
		}
	}

	return true
}

func isValid(arr []byte) bool {
	for i, b := range arr {
		if b == '.' {
			continue
		}
		if slices.Contains(arr[i+1:], b) {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	board := Deserialize[[][]byte](ReadLine(stdin))
	ans := isValidSudoku(board)

	fmt.Println("\noutput:", Serialize(ans))
}
