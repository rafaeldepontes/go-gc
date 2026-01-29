package object

// This should looks like the "Objects" from any other GC language...
type Object struct {
	ID   uint64
	Mark bool
	Refs []*Object
}
