// Created by Ming Lee Ng at 2024/04/02 13:38
// leetgo: dev
// https://leetcode.com/problems/isomorphic-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isIsomorphic(s string, t string) bool {
	mapS, mapT := [256]byte{}, [256]byte{}

	for i := 0; i < len(s); i++ {
		if mapS[s[i]] != mapT[t[i]] {
			return false
		}
		mapS[s[i]] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isIsomorphic(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
