/*

NOTE: Runs on [https://tour.golang.org/concurrency/8]

Exercise: Equivalent Binary Trees
1. Implement the Walk function.
2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted)
binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2,
3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store
the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1),
tree.New(2)) should return false.

The documentation for Tree can be found here [https://godoc.org/golang.org/x/tour/tree#Tree].

*/

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

/*
// COMMENTED DUE TO TEST21 INTERFERENCE.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}
*/
// Walks the tree, t, sending all values from tree to channel, ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether t1 and t2 contain same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}
