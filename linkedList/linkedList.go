package main

import "fmt"

// Node 节点结构体
type Node struct {
	Value int
	Next  *Node
}

// 保存链表第一个元素的全局变量
var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> empty list!")
		return 0
	}

	i := 0
	for t != nil {
		t = t.Next
		i++
	}

	return i
}

func delNode(t *Node, v int) bool {
	if t == nil {
		fmt.Println("empty list!")
		return false
	}

	if !lookupNode(t, v) {
		fmt.Println("Node does not exist!")
		return false
	}

	var pre *Node
	for t != nil {
		if t.Value == v {
			if pre == nil {
				t.Value = t.Next.Value
				t.Next = t.Next.Next
				return true
			}

			pre.Next = t.Next
			return true
		}

		pre = t
		t = t.Next
	}

	return false
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	delNode(root, 45)
	traverse(root)
	delNode(root, 1)
	traverse(root)
	delNode(root, -1)
	traverse(root)
}
