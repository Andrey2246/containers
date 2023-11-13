package containers

import "errors"

type Queue struct {
	head *sNode
	tail *sNode
}

func (s *Queue) Push(val string) { // нет циклов, все функции линейные => O(1)
	newNode := &sNode{data: val}
	if s.tail == nil {
		s.tail = newNode
		s.head = newNode
	} else {
		s.tail.next = newNode  // пишем в конец
		s.tail = newNode
	}
}

func (s *Queue) Pop() (string, error) { // нет циклов, все функции линейные => O(1)
	if s.tail == nil {
		return "", errors.New("queue is empty")
	}
	val := s.head.data
	s.head = s.head.next // берем с головы
	if s.head == nil {
		s.tail = s.head
	}
	return val, nil
}
