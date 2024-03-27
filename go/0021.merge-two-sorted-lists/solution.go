// Created by Ming Lee Ng at 2024/03/27 16:03
// leetgo: dev
// https://leetcode.com/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	head := &ListNode{}
	current := head

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			break
		}
		if list2 == nil {
			current.Next = list1
			break
		}
		val1, val2 := list1.Val, list2.Val

		if val1 < val2 {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}

	return head.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	list1 := Deserialize[*ListNode](ReadLine(stdin))
	list2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}
