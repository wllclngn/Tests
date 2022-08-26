// WEIRD, CONCURRENT DEPTH FIRST SEARCH
// "Work in Progress"

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

/*
        0
     1    2
   3  5  4  6
7 9        8 10
*/

func (node *Node) insert(data int) {

    if (node.left == nil || node.right == nil) {
        if node.left == nil && node.right == nil {
            node.left = &Node{key: data}
        } else if node.left != nil && node.right == nil {
            node.right = &Node{key: data}
        } else if (data % 2) == 0 {
            if ((node.left.key % 2) == 0) && node.left.left == nil {
                node.left.left = &Node{key: data}
            } else if ((node.left.key % 2) == 0) && node.left.right == nil {
                node.left.right = &Node{key: data}
            }
        } else if (data % 2) != 0 {
            if ((node.right.key % 2) != 0) && node.right.left == nil {
                node.right.left = &Node{key: data}
            } else if ((node.right.key % 2) != 0) && node.right.right == nil {
                node.right.right = &Node{key: data}
            }
        }
    } else {
        if (data % 2) == 0 {
            node.right.insert(data)
        } else {
            node.left.insert(data)
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

    fmt.Printf("🤑 💀 #%v LEFT: %v RIGHT: %v\n", node.key, node.left, node.right)

}

func main() {

    processors := runtime.GOMAXPROCS(runtime.NumCPU())

    start := time.Now()

    var tree Tree

    for i := 0; i <= 20; i++ {
        tree.insert(i)
    }

    wg.Add(1)

    go tree.root.DFSconcurrent()

    wg.Wait()

    elapsed := time.Since(start)
    fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, elapsed)

}
