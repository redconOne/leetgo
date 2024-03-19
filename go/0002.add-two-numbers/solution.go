// Created by Ming Lee Ng at 2024/03/19 15:48
// leetgo: dev
// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	var (
		head    *ListNode
		current *ListNode
		carry   = 0
	)

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if head == nil {
			head = &ListNode{Val: sum % 10}
			current = head
		} else {
			current.Next = &ListNode{Val: sum % 10}
			current = current.Next
		}

		carry = sum / 10
	}

	return head
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
