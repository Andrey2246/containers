package containers

import "errors"

type sNode struct {
	next *sNode
	data string
}
type Stack struct {
	head *sNode
}

func (s *Stack) Push(val string) {
	newNode := &sNode{data: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Pop() (string, error) {
	if s.head == nil {
		return "", errors.New("stack is empty")
	}
	val := s.head.data
	s.head = s.head.next
	return val, nil
}
