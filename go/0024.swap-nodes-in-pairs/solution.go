// Created by Ming Lee Ng at 2024/03/28 08:43
// leetgo: dev
// https://leetcode.com/problems/swap-nodes-in-pairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func swapPairs(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head

	return head
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := swapPairs(head)

	fmt.Println("\noutput:", Serialize(ans))
}
