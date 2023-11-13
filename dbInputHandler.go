package containers //hehe

import (
	"strconv"
	"sync"
)

type DataBase struct {
	Stack     map[string]*Stack
	HashMap     map[string]*HashMap
	Queue    map[string]*Queue
	Set   map[string]*Set
	BST    map[string]*Bst
	Array    map[string]*Arr
	Mutex sync.Mutex
}

func (db *DataBase) Execute(commands *Arr, password string) string {
	command := commands.Get(0)
	key := commands.Get(1)
	val := commands.Get(2)
	if command != "exit" && key == ""{
		return "no key"
	}
	if command == "" {
		return "no command"
	}
	switch command[0] {                        //создаем контейнер под используемый логин, если еще не создали (в го нет конструкторов)
	case 'H':
		{
			if db.HashMap[password] == nil {
				db.HashMap[password] = new(HashMap)
			}
		}
	case 'S':
		{
			if command[1] != 'P' && db.Set[password] == nil { // SP => SPUSH || SPOP
				db.Set[password] = new(Set)
			} else if db.Stack[password] == nil {
				db.Stack[password] = new(Stack)
			}
		}
	case 'B':
		{
			if db.BST[password] == nil {
				db.BST[password] = new(Bst)
			}
		}
	case 'Q':
		if db.Queue[password] == nil {
			db.Queue[password] = new(Queue)
		}
	case 'A':
		{
			if db.Array[password] == nil {
				db.Array[password] = new(Arr)
			}
		}
	}
	switch command {                                   //здесь есть команды из лабы 1, ее еще пишу
	case "HSET":
		{
			if val == "" {
				return "no value"
			}
			err := db.HashMap[password].Set(key, val)
			if err != nil {
				return err.Error()
			}
		}
	case "HGET":
		{
			st, err := db.HashMap[password].Get(key)
			if err != nil {
				return err.Error()
			} else {
				return ("Value on key" + key + ":" + st)
			}
		}
	case "HDEL":
		{
			err := db.HashMap[password].Del(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SPUSH":
		{
			db.Stack[password].Push(key)
		}
	case "SPOP":
		{
			st, err := db.Stack[password].Pop()
			if err != nil {
				return err.Error()
			} else {
				return st
			}
		}
	case "QPUSH":
		{
			db.Queue[password].Push(key)
		}
	case "QPOP":
		{
			st, err := db.Queue[password].Pop()
			if err != nil {
				return err.Error()
			} else {
				return st
			}
		}
	case "BADD":
		{
			err := db.BST[password].Add(key)
			if err != nil {
				return err.Error()
			}
		}
	case "BPRINT":
		{
			return db.BST[password].Print()
		}
	case "BDEL":
		{
			err := db.BST[password].Del(key)
			if err != nil {
				return err.Error()
			}
		}
	case "BISMEM":
		{
			if db.BST[password].IsMem(key) {
				return ("Your tree contains value\"" + key + "\"")
			} else {
				return ("Your tree does not contain value\"" + key + "\"")
			}
		}
	case "ASET":
		{
			keyInt, _ := strconv.Atoi(key)
			db.Array[password].Set(keyInt, val)
		}
	case "AGET":
		{
			keyI, err := strconv.Atoi(key)
			if err != nil {
				return "key to array is an integer number, not" + key
			} else {
				db.Array[password].Get(keyI)
			}
		}
	case "SADD":
		{
			err := db.Set[password].ADD(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SREM":
		{
			err := db.Set[password].Rem(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SISMEM":
		{
			ans := db.Set[password].IsMem(key)
			if ans {
				return "Your set contains value \"" + key + "\""
			}
			return "Your set does not contain value \"" + key + "\""
		}
	case "exit":
		{
			return "exit"
		}
	default:
		return "no command: " + command
	}
	return "OK"
}
