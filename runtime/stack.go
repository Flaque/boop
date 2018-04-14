package runtime

import s "github.com/golang-collections/collections/stack"

type Stack struct {
	lower *s.Stack
}

func NewStack() Stack {
	st := s.New()
	return Stack{lower: st}
}

func (st *Stack) Peek() Frame {
	val, _ := st.lower.Peek().(Frame)
	return val
}

func (st *Stack) Push(f Frame) {
	st.lower.Push(f)
}

func (st *Stack) Pop() Frame {
	val, _ := st.lower.Pop().(Frame)
	return val
}

func (st *Stack) Len() int {
	return st.lower.Len()
}
