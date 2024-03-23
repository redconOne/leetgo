// Created by Ming Lee Ng at 2024/03/22 17:23
// leetgo: dev
// https://leetcode.com/problems/palindrome-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverser(slow.Next)

	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}

func reverser(head *ListNode) *ListNode {
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
	ans := isPalindrome(head)

	fmt.Println("\noutput:", Serialize(ans))
}
