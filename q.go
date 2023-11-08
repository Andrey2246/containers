package containers

import "errors"

type Queue struct {
	head *sNode
	tail *sNode
}

func (s *Queue) Push(val string) {
	newNode := &sNode{data: val}
	if s.tail == nil {
		s.tail = newNode
		s.head = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
}

func (s *Queue) Pop() (string, error) {
	if s.tail == nil {
		return "", errors.New("queue is empty")
	}
	val := s.head.data
	s.head = s.head.next
	if s.head == nil {
		s.tail = s.head
	}
	return val, nil
}
