// Created by Ming Lee Ng at 2024/03/22 02:37
// leetgo: dev
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type permStruct struct {
	Digits  []rune
	Lookup  map[rune]string
	Results []string
}

func letterCombinations(digits string) (ans []string) {
	lookup := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	data := &permStruct{
		Digits:  []rune(digits),
		Lookup:  lookup,
		Results: []string{},
	}

	if digits == "" {
		return data.Results
	}

	getCombinations(0, "", data)
	return data.Results
}

func getCombinations(start int, combo string, data *permStruct) {
	if start >= len(data.Digits) {
		data.Results = append(data.Results, combo)
		return
	}

	ch := data.Digits[start]
	values := data.Lookup[ch]

	for _, c := range values {
		temp := combo + string(c)
		getCombinations(start+1, temp, data)
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	digits := Deserialize[string](ReadLine(stdin))
	ans := letterCombinations(digits)

	fmt.Println("\noutput:", Serialize(ans))
}
