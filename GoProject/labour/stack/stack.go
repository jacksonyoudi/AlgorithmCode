package stack

import (
	"container/list"
)

// stack struct
type Stack struct {
	list *list.List
}

// get a stack
func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

// stack method
// push  list back end
func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

// pop  back element
func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil

}

func (s *Stack) Peek() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
