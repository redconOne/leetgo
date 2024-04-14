// Created by Ming Lee Ng at 2024/04/14 17:16
// leetgo: dev
// https://leetcode.com/problems/sum-of-left-leaves/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	return helper(root, false)
}

func helper(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}

	return helper(node.Left, true) + helper(node.Right, false)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := sumOfLeftLeaves(root)

	fmt.Println("\noutput:", Serialize(ans))
}
