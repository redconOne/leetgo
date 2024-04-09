// Created by Ming Lee Ng at 2024/04/08 20:26
// leetgo: dev
// https://leetcode.com/problems/min-stack/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MinStack struct {
	vals   []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{
		vals:   []int{},
		minVal: 1<<63 - 1,
	}
}

func (m *MinStack) Push(val int) {
	m.vals = append(m.vals, val)
	m.minVal = min(m.minVal, val)
}

func (m *MinStack) Pop() {
	top := m.Top()

	m.vals = m.vals[:len(m.vals)-1]

	if top == m.minVal && len(m.vals) > 0 {
		newMin := m.vals[0]

		for _, num := range m.vals {
			if newMin > num {
				newMin = num
			}
		}

		m.minVal = newMin
	}

	if len(m.vals) == 0 {
		m.minVal = 1<<63 - 1
	}
}

func (m *MinStack) Top() (ans int) {
	return m.vals[len(m.vals)-1]
}

func (m *MinStack) GetMin() (ans int) {
	return m.minVal
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "push":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.Push(val)
			output = append(output, "null")
		case "pop":
			obj.Pop()
			output = append(output, "null")
		case "top":
			ans := Serialize(obj.Top())
			output = append(output, ans)
		case "getMin":
			ans := Serialize(obj.GetMin())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
