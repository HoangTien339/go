package main

import (
	"./lib"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
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
	c1, c2 := make(chan int), make(chan int)
	go func() {
		defer close(c1)
		Walk(t1, c1)
	}()
	go func() {
		defer close(c2)
		Walk(t2, c2)
	}()
	for {
		fmt.Println("-----------------------------------------")
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		fmt.Printf("v1, ok1 = %v, %v | v2, ok2 = %v, %v\n", v1, ok1, v2, ok2)
		if ok1 != ok2 || v1 != v2 {
			fmt.Println("Not same !")
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	fmt.Println("Same !")
	return true
}

func main() {
	k := 1
	t1 := tree.New(k)
	t2 := tree.New(k)
	Same(t1, t2)
	t2 = tree.Insert(t2, 5)
	Same(t1, t2)
}
