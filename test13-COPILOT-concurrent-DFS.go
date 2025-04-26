package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

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
	if node.left == nil || node.right == nil {
		if node.left == nil && node.right == nil {
			node.left = &Node{key: data}
		} else if node.left != nil && node.right == nil {
			node.right = &Node{key: data}
		}
	} else {
		if (data % 2) == 0 {
			if ((node.left.key % 2) == 0) && node.left.left == nil {
				node.left.left = &Node{key: data}
			} else if ((node.left.key % 2) == 0) && node.left.right == nil {
				node.left.right = &Node{key: data}
			} else {
				node.right.insert(data)
			}
		} else {
			if ((node.right.key % 2) != 0) && node.right.left == nil {
				node.right.left = &Node{key: data}
			} else if ((node.right.key % 2) != 0) && node.right.right == nil {
				node.right.right = &Node{key: data}
			} else {
				node.left.insert(data)
			}
		}
	}
}

// Recursive DFS function with bounded concurrency
func (node *Node) DFSconcurrentRecursive(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done() // Ensure Done is called when this function exits

	if node == nil {
		return
	}

	// Print the current node
	fmt.Printf("Visiting Node: %d\n", node.key)

	// Launch recursive goroutines for left and right children with bounded concurrency
	if node.left != nil {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot in the semaphore
		go func(left *Node) {
			defer func() { <-semaphore }() // Release the slot in the semaphore
			left.DFSconcurrentRecursive(wg, semaphore)
		}(node.left)
	}

	if node.right != nil {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot in the semaphore
		go func(right *Node) {
			defer func() { <-semaphore }() // Release the slot in the semaphore
			right.DFSconcurrentRecursive(wg, semaphore)
		}(node.right)
	}
}

func main() {
	// Set the maximum number of CPUs to use
	processors := runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	// Create and populate the tree
	var tree Tree
	for i := 0; i <= 20; i++ {
		tree.insert(i)
	}

	// WaitGroup to track all goroutines
	var wg sync.WaitGroup

	// Semaphore to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, 10) // Limit to 4 concurrent goroutines

	// Start the DFS from the root
	wg.Add(1)
	go tree.root.DFSconcurrentRecursive(&wg, semaphore)

	// Wait for all goroutines to complete
	wg.Wait()

	// Measure and print execution time
	elapsed := time.Since(start)
	fmt.Printf("\nProcessors: %d, Time elapsed: %v\n", processors, elapsed)
}
