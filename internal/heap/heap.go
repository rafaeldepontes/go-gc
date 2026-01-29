package heap

import "github.com/rafaeldepontes/go-gc/internal/object"

type Heap struct {
	objs   map[uint64]*object.Object
	nextID uint64
}

// I would try to build malloc from C, but the void pointer thing is not
// my style...
func (h *Heap) Alloc(refs ...*object.Object) *object.Object {
	obj := &object.Object{
		ID:   h.nextID,
		Refs: refs,
	}

	h.objs[obj.ID] = obj
	h.nextID++
	return obj
}

func (h *Heap) Sweep() {
	for id, obj := range h.objs {
		if !obj.Mark {
			delete(h.objs, id)
			continue
		}
		obj.Mark = false
	}
}
