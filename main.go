package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	var v1, v2 string
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	for v := range ch1 {
		v1 += fmt.Sprintf("%d", v)
	}
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	for v := range ch2 {
		v2 += fmt.Sprintf("%d", v)
	}
	if v1 == v2 {
		return true
	}
	return false
}

func main() {
	result1 := Same(tree.New(1), tree.New(1))
	fmt.Println(result1)
}
