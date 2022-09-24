package main

import (
	"github.com/MrHenri/logbook/goroutine/binarytrees"
	"golang.org/x/tour/tree"
)

func main() {
	binarytrees.Same(tree.New(4), tree.New(4))
	binarytrees.Same(tree.New(3), tree.New(4))
}
