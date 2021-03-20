// WEIRD, CONCURRENT DFS

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Node struct {
	Data interface{}

	Sleep time.Duration

	Left *Node

	Right *Node
}

func NewNode(data interface{}) *Node {

	node := new(Node)

	node.Data = data
	node.Left = nil
	node.Right = nil

	return node

}

func (n *Node) DFSParallel() {

	defer wg.Done()

	if n == nil {
		return
	}

	wg.Add(3)

	go n.ProcessNodeParallel()

	go n.Left.DFSParallel()

	go n.Right.DFSParallel()

}

func (n *Node) ProcessNodeParallel() {

	defer wg.Done()

	fmt.Printf("#%v 🚀\n", n.Data)

}

func main() {

	start := time.Now()

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	wg.Add(1)

	go root.DFSParallel()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, elapsed)
}
