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

type Node struct {
	key   int
	left  *Node
	right *Node
}

// Tree
func (tree *Tree) insert(data int) {
	if tree.root == nil {
		tree.root = &Node{key: data}
	} else {
		tree.root.insert(data)
	}
}

// Node
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

func (node *Node) DFSParallel() {

	defer wg.Done()

	if node == nil {
		return
	}

	wg.Add(2)

	go node.left.DFSParallel()

	go node.right.DFSParallel()

	//time.Sleep(time.Millisecond * 200)

	fmt.Printf("#%v 🤑💀\n", node.key)

}

func main() {

	start := time.Now()

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	var tree Tree

	for i := 1; i <= 10; i++ {
		tree.insert(i)
	}

	wg.Add(1)

	go tree.root.DFSParallel()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, elapsed)

}
