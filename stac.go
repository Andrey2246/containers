package containers

import "errors"

type sNode struct {
	next *sNode
	data string
}
type Stack struct {
	head *sNode
}

func (s *Stack) Push(val string) { // нет циклов, все функции линейные => O(1)
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

func (s *Stack) CheckBraces(str string) (bool, error) {
	for i := range str {
		l := string(str[i])
		if l != "(" && l != ")" && l != "{" && l != "}" && l != "[" && l != "]" {
			return false, errors.New("not a brace found")
		}
		if l == ")" || l == "}" || l == "]" {
			b, err := s.Pop()
			if err != nil {
				return false, nil
			}
			if (l != ")" || b != "(") && (l != "}" || b != "{") && (l != "]" || b != "[") {
				return false, nil
			}
		}
		s.Push(l)
	}
	return true, nil
}
