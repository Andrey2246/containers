package containers

import "errors"

type sNode struct {
	next *sNode
	data string
}
type Stack struct {
	head *sNode
}

func (s *Stack) Push(val string) {  // нет циклов, все функции линейные => O(1)
	newNode := &sNode{data: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Pop() (string, error) { // нет циклов, все функции линейные => O(1)
	if s.head == nil {
		return "", errors.New("stack is empty")
	}
	val := s.head.data
	s.head = s.head.next
	return val, nil
}
