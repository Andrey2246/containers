package containers

import (
	"errors"
	"strconv"
)

type Bst struct {
	val   string
	left  *Bst
	right *Bst
}

func (tree *Bst) Add(val string) error {
	if tree.val == val {
		return errors.New("elem already exists")
	}
	if tree.val == "" {
		tree.val = val
		return nil
	}
	if val < tree.val {
		if tree.left == nil {
			tree.left = &Bst{val: val}
			return nil
		}
		return tree.left.Add(val)
	}
	if tree.val < val {
		if tree.right == nil {
			tree.right = &Bst{val: val}
			return nil
		}
		return tree.right.Add(val)
	}
	return errors.New("cannot add new element")
}

func (tree *Bst) Del(val string) error {
	if !tree.IsMem(val) {
		return errors.New("no such element")
	}
	if val < tree.val {
		//t, err :=
		tree.left.Del(val)
		return nil
	} else if val > tree.val {
		tree.right.Del(val)
		return nil
	} else {
		cur := tree
		if tree.left == nil && tree.right == nil {
			tree.val = ""
			tree = nil
		} else if tree.left == nil {
			*tree = *tree.right
		} else if tree.right == nil {
			*tree = *tree.left
		} else {
			cur := cur.right
			for cur != nil && cur.left != nil && cur.right != nil {
				if cur.left != nil {
					*cur = *cur.left
				} else if cur.right != nil {
					*cur = *cur.right
				} else {
					break
				}
			}
			tree.val = cur.val
			cur = nil
		}
		return nil
	}
}

func (tree *Bst) IsMem(val string) bool {
	return tree.val == val || tree.left != nil && tree.left.IsMem(val) || tree.right != nil && tree.right.IsMem(val)
}

func (tree *Bst) Print(params ...string) string {
	if params == nil {
		params = make([]string, 2)
		params[0] = ""
		params[1] = "0"
	}
	str := params[0]
	level, _ := strconv.Atoi(params[1])
	if tree != nil {
		str = tree.left.Print(str, strconv.Itoa(level+1))
		for i := 0; i < level; i++ {
			str += "\t"
		}
		str += tree.val
		str += "\n"
		str = tree.right.Print(str, strconv.Itoa(level+1))
	}
	return str
}
