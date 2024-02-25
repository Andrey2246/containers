package containers

import (
	"errors"
	"sort"
	"strconv"
)

type Set struct { // также как хеш-таблица, но клуч = значние
	table [512]*setNode
}

type setNode struct {
	val string
}

func (set *Set) ADD(key string) error { // нет циклов, все функции линейные => O(1)
	p := new(setNode) //обработка коллизии в сложность алгоритма не включаетcя
	p.val = key
	hash, err := calcHash(key, len(set.table)) //хеш-функция в файле hash.go
	if err != nil {
		return errors.New("unacceptable key")
	}
	if set.table[hash] == nil {
		set.table[hash] = p
		return nil
	}
	if set.table[hash].val == key {
		return errors.New("this elem already exists")
	}
	for i := (hash + 1) % len(set.table); i != hash; i = (i + 1) % len(set.table) {
		if set.table[i] == nil {
			set.table[i] = p
			return nil
		}
		if set.table[hash].val == key {
			return errors.New("this elem already exists")
		}
	}
	return errors.New("table is full")
}

func (set *Set) Rem(key string) error { // нет циклов, все функции линейные => O(1)
	hash, err := calcHash(key, len(set.table)) //обработка коллизии в сложность алгоритма не включаетcя
	if err != nil {
		return errors.New("unacceptable key")
	}
	if set.table[hash] == nil {
		return errors.New("nothing to delete")
	}
	if set.table[hash] != nil && set.table[hash].val == key {
		set.table[hash] = nil
		return nil
	}
	for i := hash + 1; i != hash; i = (i + 1) % len(set.table) {
		if set.table[i] != nil && set.table[i].val == key {
			set.table[i] = nil
			return nil
		}
	}
	return errors.New("no such key")
}

func (set *Set) IsMem(key string) bool { // нет циклов, все функции линейные => O(1)
	hash, err := calcHash(key, len(set.table)) //обработка коллизии в сложность алгоритма не включаетcя
	if err != nil {
		return false
	}
	if set.table[hash] != nil && set.table[hash].val == key {
		return true
	}
	for i := (hash + 1) % len(set.table); i != hash; i = (i + 1) % len(set.table) {
		if set.table[i] != nil && set.table[i].val == key {
			return true
		}
	}
	return false
}

func (set *Set) GetAllSet() ([]int, error) {
	ans := make([]int, 1)
	for i := range set.table {
		if set.table[i] != nil {
			temp, err := strconv.Atoi(set.table[i].val)
			if err != nil {
				return nil, errors.New("not int found")
			}
			ans = append(ans, temp)
		}
	}
	return ans, nil
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (set *Set) EqualDivide(div int) ([]*Set, error) {
	arr, err := set.GetAllSet()
	sort.Slice(arr, func(i, j int) bool { return j < i })
	if err != nil {
		return nil, err
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%div != 0 {
		return nil, errors.New("sum mod div != 0")
	}
	n := sum / div
	ans := make([][]int, n)
	l := len(arr)
	for i := 0; arr != nil; i++ {

		ans[i] = append(ans[i], arr[0]) //take elem
		arr = remove(arr, 0)
		l -= 1

		tempS := sum // make temp sum
		tempS -= arr[0]
		if tempS < 0 { //check temp sum
			return nil, errors.New("found elem > div")
		}
		if tempS == 0 {
			continue
		}
		for j := range arr[:] {
			if arr[j] > tempS && j != len(arr)-1 { //skip elem if > temp sum
				if j == len(arr)-1 { //err if no needed elem
					return nil, errors.New("could not divide equally")
				} else {
					continue
				}
			}
			ans[i] = append(ans[i], arr[j]) //take elem
			arr = remove(arr, j)
			l -= 1

			tempS -= arr[0]
			if tempS == 0 {
				break
			}
		}
	}

}
