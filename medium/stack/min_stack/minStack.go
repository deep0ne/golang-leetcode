// https://leetcode.com/problems/min-stack/description/?envType=study-plan-v2&envId=top-interview-150
package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, val)
	} else {
		if val <= m.minStack[len(m.minStack)-1] {
			m.minStack = append(m.minStack, val)
		}
	}
}

func (m *MinStack) Pop() {
	val := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if val == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
