// Created by Ming Lee Ng at 2024/03/22 18:08
// leetgo: dev
// https://leetcode.com/problems/reorder-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reorderList(head *ListNode) {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	secondList := slow.Next
	slow.Next = nil

	secondList = reverse(secondList)

	a, b := head, secondList

	for b != nil {
		aNext := a.Next
		bNext := b.Next
		a.Next = b
		b.Next = aNext
		a = aNext
		b = bNext
	}
}

func reverse(head *ListNode) *ListNode {
	var (
		prev *ListNode
		next *ListNode
	)

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	reorderList(head)
	ans := head

	fmt.Println("\noutput:", Serialize(ans))
}
