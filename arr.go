package containers

type aNode struct {
	next *aNode
	key  int
	val  string
}
type Arr struct {
	head *aNode
	len  int
}

func (a *Arr) Set(key int, val string) {
	if a.head == nil {
		(*a).head = &aNode{key: 0, val: ""}
	}
	node := a.head
	i := 0
	newNode := &aNode{key: key, val: val}
	for i < key {
		if (*node).next == nil {
			(*node).next = &aNode{key: i, val: ""}
		}
		node = node.next
		i += 1
	}
	*node = *newNode
	a.len = max(i, a.len)
}

func (a *Arr) Get(key int) string {
	node := a.head
	i := 0
	for i < key {
		if node.next == nil {
			return ""
		}
		node = node.next
		i += 1
	}
	return node.val
}
