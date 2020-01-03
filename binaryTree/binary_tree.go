package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// 非平衡二叉树
func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}

// 递归访问二叉树上的所有节点
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Value, " ")
	traverse(t.Right)
}

// 向树中填充随机的数值
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

// insert() 函数使用 if 语句做了很多重要的事。第一个 if 语句检查要操作的树是否为空。
// 如果是空树，那么通过 &Tree{nil, v, nil} 创建的新节点将成为该树的根节点。
// 第二个 if 语句判断二叉树上是否已经存在将要插入的值。
// 如果值已经存在，那么函数将什么也不做然后返回。
// 第三个 if 语句判断对于当前节点，被插入的值是在节点的左侧还是右侧，然后执行相应的操作。
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}
