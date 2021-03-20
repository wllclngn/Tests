// WEIRD, CONCURRENT DFS

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
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *Node) DFSParallel() {

	defer wg.Done()

	if n == nil {
		return
	}

	wg.Add(3)

	go n.ProcessNodeParallel()

	go n.left.DFSParallel()

	go n.right.DFSParallel()

}

func (n *Node) ProcessNodeParallel() {

	defer wg.Done()

	fmt.Printf("#%v 🚀\n", n.key)

}

func main() {

	start := time.Now()

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	var t Tree

	for i := 1; i <= 10; i++ {
		t.insert(i)
	}

	wg.Add(1)

	go t.root.DFSParallel()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, elapsed)
}
