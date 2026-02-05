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

	println("Before GC...")
	fmt.Printf("\tObject's:\n\t\t-A: %v\n\t\t-B: %v\n\t\t-C: %v\n\n",
		vm.Heap.Objs[a.ID],
		vm.Heap.Objs[b.ID],
		vm.Heap.Objs[c.ID],
	)
	vm.GC()
	println("After GC...")
	fmt.Printf(
		"\tHeap's Object:\n\t\t-A: %v\n\t\t-B: %v\n\t\t-C: %v\n",
		vm.Heap.Objs[a.ID],
		vm.Heap.Objs[b.ID],
		vm.Heap.Objs[c.ID],
	)

}
