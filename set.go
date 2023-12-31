package containers

import "errors"

type Set struct {           // также как хеш-таблица, но клуч = значние 
	table [512]*setNode
}

type setNode struct {
	val string
}

func (set *Set) ADD(key string) error { // нет циклов, все функции линейные => O(1)
	p := new(setNode)					//обработка коллизии в сложность алгоритма не включаетcя
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
/*
func (set *Set) Get(key string) (string, error) { // нет циклов, все функции линейные => O(1)
	hash, err := calcHash(key, len(set.table))
	if err != nil {
		return "", errors.New("unacceptable key")
	}
	if set.table[hash] != nil && set.table[hash].val == key {
		return set.table[hash].val, nil
	}
	for i := (hash + 1) % len(set.table); i != hash; i = (i + 1) % len(set.table) {
		if set.table[i] != nil && set.table[i].val == key {
			return set.table[i].val, nil
		}
	}
	return "", errors.New("no such key")
}
*/
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
