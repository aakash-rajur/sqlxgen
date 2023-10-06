package linked_list

type Node[T any] struct {
	Value *T
	Prev  *Node[T]
	Next  *Node[T]
}
