// Created by Ming Lee Ng at 2024/04/10 15:42
// leetgo: dev
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkIfPangram(sentence string) bool {
	dict := make(map[rune]bool)

	for _, r := range sentence {
		dict[r] = true
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, c := range alphabet {
		if found := dict[c]; !found {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	sentence := Deserialize[string](ReadLine(stdin))
	ans := checkIfPangram(sentence)

	fmt.Println("\noutput:", Serialize(ans))
}
