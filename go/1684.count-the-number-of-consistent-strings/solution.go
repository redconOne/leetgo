// Created by Ming Lee Ng at 2024/04/11 17:22
// leetgo: dev
// https://leetcode.com/problems/count-the-number-of-consistent-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countConsistentStrings(allowed string, words []string) (ans int) {
	count := 0
	allowedSet := NewSet()

	for _, c := range allowed {
		allowedSet.addChars(c)
	}

	for _, word := range words {
		wordSet := NewSet()

		for _, c := range word {
			wordSet.addChars(c)
		}

		if allowedSet.compare(wordSet) {
			count++
		}
	}

	return count
}

type Set struct {
	items map[rune]bool
}

func NewSet() *Set {
	return &Set{
		items: make(map[rune]bool),
	}
}

func (s *Set) addChars(chars ...rune) {
	for _, c := range chars {
		s.items[c] = true
	}
}

func (s *Set) size() int {
	return len(s.items)
}

func (s1 *Set) compare(s2 *Set) bool {
	for c := range s2.items {
		if !s1.items[c] {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	allowed := Deserialize[string](ReadLine(stdin))
	words := Deserialize[[]string](ReadLine(stdin))
	ans := countConsistentStrings(allowed, words)

	fmt.Println("\noutput:", Serialize(ans))
}
