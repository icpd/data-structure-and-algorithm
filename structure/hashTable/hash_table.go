package main

// https://wskdsgcf.gitbook.io/mastering-go-zh-cn/05.0/05.4/05.4.1 mastering go 哈希表

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 15 // SIZE 常量表示哈希表中桶的数量。

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

// 向哈希表中插入元素
func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

// 输出哈希表所有的值
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if v, ok := hash.Table[k]; ok && v != nil {
			for v != nil {
				fmt.Printf("%d -> ", v.Value)
				v = v.Next
			}
		}
		fmt.Println()
	}
}

// 检查哈希表中是否已经存在某个给定的元素
func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Numbder of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)

	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(300)
		if !lookup(hash, v) {
			fmt.Println(v, "is not in the hash table!")
		}
	}
}
