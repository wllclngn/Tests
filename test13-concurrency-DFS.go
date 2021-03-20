// WEIRD, CONCURRENT DEPTH FIRST SEARCH

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Tree struct {
	root *Node
}

func (tree *Tree) insert(data int) {
	if tree.root == nil {
		tree.root = &Node{key: data}
	} else {
		tree.root.insert(data)
	}
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (node *Node) insert(data int) {
	if data <= node.key {
		if node.left == nil {
			node.left = &Node{key: data}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{key: data}
		} else {
			node.right.insert(data)
		}
	}
}

func (node *Node) DFSconcurrent() {

	defer wg.Done()

	if node == nil {
		return
	}

	wg.Add(2)

	go node.left.DFSconcurrent()

	go node.right.DFSconcurrent()

	fmt.Printf("🤑 #%v 💀\n", node.key)

}

func main() {

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var tree Tree

	for i := 0; i <= 10; i++ {
		tree.insert(i)
	}

	wg.Add(1)

	go tree.root.DFSconcurrent()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, elapsed)

}
