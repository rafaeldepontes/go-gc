package gc

import (
	"sync"

	"github.com/rafaeldepontes/go-gc/internal/heap"
	"github.com/rafaeldepontes/go-gc/internal/object"
)

var wg sync.WaitGroup

// The logic here is pretty simple, if the logic is inside of the root
// it should be still a live after the gc calls...
type VM struct {
	Heap  *heap.Heap
	Roots []*object.Object
}

func NewVM() *VM {
	return &VM{
		Heap:  heap.NewHeap(),
		Roots: make([]*object.Object, 0, 15),
	}
}

// This is called tri-color marking, it goes from white to black blazing
// fast!!
func (vm *VM) markAll() {
	for _, root := range vm.Roots {
		wg.Add(1)
		go mark(root)
	}
}

func (vm *VM) GC() {
	wg.Add(1)
	vm.markAll()
	wg.Wait()
	vm.Heap.Sweep()
}

// This is done recursively, not sure if it's the best choice... but is
// the one I'm going for noow...
func mark(obj *object.Object) {
	defer wg.Done()
	if obj.Mark {
		return
	}
	obj.Mark = true
	for _, ref := range obj.Refs {
		mark(ref)
	}
}
