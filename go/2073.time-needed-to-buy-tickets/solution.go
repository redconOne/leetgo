// Created by Ming Lee Ng at 2024/04/09 15:15
// leetgo: dev
// https://leetcode.com/problems/time-needed-to-buy-tickets/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func timeRequiredToBuy(tickets []int, k int) (ans int) {
	count, ticketLen := 0, len(tickets)

	if ticketLen == 0 {
		return count
	}

	for i := 0; tickets[k] > 0; i++ {
		if i >= ticketLen {
			i = 0
		}
		if tickets[i] <= 0 {
			continue
		}

		tickets[i]--
		count++
	}

	return count
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tickets := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := timeRequiredToBuy(tickets, k)

	fmt.Println("\noutput:", Serialize(ans))
}
