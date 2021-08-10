package main

import "fmt"

// Node 双向链表的节点
type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("empty list!")
		return
	}

	for t.Next != nil {
		t = t.Next
	}

	for t.Previous != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Previous
	}

	fmt.Printf("%d ->\n", t.Value)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
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

func delNode(t *Node, v int) {
	if t == nil {
		return
	}

	for t != nil {
		if t.Value == v {
			if t.Previous != nil {
				t.Previous.Next = t.Next
				return
			}

			t.Value = t.Next.Value
			t.Previous = nil
			t.Next = t.Next.Next
			return
		}
		t = t.Next
	}
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
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
	delNode(root, -1)
	traverse(root)
	delNode(root, 1)
	traverse(root)
}
