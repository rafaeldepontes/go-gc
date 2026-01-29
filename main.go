package main

import (
	"fmt"

	gc "github.com/rafaeldepontes/go-gc/internal/gc"
	obj "github.com/rafaeldepontes/go-gc/internal/object"
)

func main() {
	vm := gc.NewVM()

	a := vm.Heap.Alloc()
	b := vm.Heap.Alloc(a)
	c := vm.Heap.Alloc()

	vm.Roots = []*obj.Object{b}

	vm.GC()

	fmt.Println(a.ID, b.ID, c.ID)
}
