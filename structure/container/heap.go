package main

// https://wskdsgcf.gitbook.io/mastering-go-zh-cn/05.0/05.9/05.9.1
import (
	"container/heap"
	"fmt"
)

type heapV []int

func (n *heapV) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	tmp := old[0 : len(old)-1]
	*n = tmp
	return x
}

func (n *heapV) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n heapV) Len() int {
	return len(n)
}

func (n heapV) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapV) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapV{11, 22, 1, -100}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(-102)
	myHeap.Push(2)
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)

	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)

	fmt.Println(myHeap.Pop())
	heap.Init(myHeap)
	fmt.Println(myHeap.Pop())
	fmt.Println(myHeap)
}
