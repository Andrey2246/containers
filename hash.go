package containers

import "errors"

type HashMap struct {
	table [512]*Pair
}

type Pair struct {
	key   string
	value string
}

func calcHash(key string, size int) (int, error) {        // хэш - сумма utf-8 кодов символов строки
	if len(key) == 0 {									  // хэш-функция тоже не учитывается при подсчете сложности алгоритма
		return 0, errors.New("no value")
	}
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % size, nil
}

func (hmap *HashMap) Set(key string, value string) error {
	p := new(Pair)										   //обработка коллизии в сложность алгоритма не включаетcя
	p.key = key
	p.value = value
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		hmap.table[hash] = p
		return nil
	}
	if hmap.table[hash].key == key {
		return errors.New("this key already exists")
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {   //если ячейка с данным хеэем занята, берем ближайшую свободную справа по кругу
		if hmap.table[i] == nil {
			hmap.table[i] = p
			return nil
		}
		if hmap.table[hash].key == key {
			return errors.New("this key already exists")
		}
	}
	return errors.New("table is full")
}

func (hmap *HashMap) Get(key string) (string, error) { // нет циклов, все функции линейные => O(1)
	hash, err := calcHash(key, len(hmap.table))       //обработка коллизии в сложность алгоритма не включаетcя
	if err != nil {
		return "", errors.New("unacceptable key")
	}
	if hmap.table[hash] != nil && hmap.table[hash].key == key {
		return hmap.table[hash].value, nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) { // так как записывали по кругу, проверять тоже надо по кругу
		if hmap.table[i] != nil && hmap.table[i].key == key {
			return hmap.table[i].value, nil
		}
	}
	return "", errors.New("no such key")
}

func (hmap *HashMap) Del(key string) error { // нет циклов, все функции линейные => O(1)
	hash, err := calcHash(key, len(hmap.table)) //обработка коллизии в сложность алгоритма не включаетcя
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		return errors.New("nothing to delete")
	}
	if hmap.table[hash] != nil && hmap.table[hash].key == key {
		hmap.table[hash] = nil
		return nil
	}
	for i := hash + 1; i != hash; i = (i + 1) % len(hmap.table) { // так как записывали по кругу, проверять тоже надо по кругу
		if hmap.table[i] != nil && hmap.table[i].key == key {
			hmap.table[i] = nil
			return nil
		}
	}
	return errors.New("no such key")
}
