// Created by Ming Lee Ng at 2024/04/06 13:12
// leetgo: dev
// https://leetcode.com/problems/water-bottles-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxBottlesDrunk(numBottles int, numExchange int) (ans int) {
	drank, empty := 0, 0

	for numBottles > 0 {
		drank += numBottles
		empty += numBottles
		numBottles = 0

		for empty >= numExchange {
			empty -= numExchange
			numBottles++
			numExchange++
		}
	}
	return drank
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numBottles := Deserialize[int](ReadLine(stdin))
	numExchange := Deserialize[int](ReadLine(stdin))
	ans := maxBottlesDrunk(numBottles, numExchange)

	fmt.Println("\noutput:", Serialize(ans))
}
