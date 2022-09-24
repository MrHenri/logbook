package binarytrees

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
	var list1 []int
	var list2 []int
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan int)
	finishJob := 0

	go func() {
		Walk(t1, ch1)
		finishJob++
		quit <- finishJob
	}()

	go func() {
		Walk(t2, ch2)
		finishJob++
		quit <- finishJob
	}()

	for {
		select {
		case i := <-ch1:
			list1 = append(list1, i)
		case i := <-ch2:
			list2 = append(list2, i)
		case j := <-quit:
			if j == 2 {
				fmt.Println(list1)
				fmt.Println(list2)
				return testEq(list1, list2)
			}
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println(false)
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			fmt.Println(false)
			return false
		}
	}
	fmt.Println(true)
	return true
}
