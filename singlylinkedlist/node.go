package singlylinkedlist

type node[T comparable] struct {
	next  *node[T]
	value T
}

func NewNode[T comparable](next *node[T], value T) *node[T] {
	newNode := node[T]{next: next, value: value}
	return &newNode
}
