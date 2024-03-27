// Created by Ming Lee Ng at 2024/03/27 15:01
// leetgo: dev
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) (ans *ListNode) {
	left, right := head, head

	for distance := 0; distance < n; distance++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right != nil && right.Next != nil {
		left, right = left.Next, right.Next
	}

	left.Next = left.Next.Next
	return head
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := removeNthFromEnd(head, n)

	fmt.Println("\noutput:", Serialize(ans))
}
