package singlylinkedlist

import "fmt"

var length int = 0

type SinglyLinkedList[T comparable] struct {
	head, tail *node[T]
}

func NewSinglyLinkListed[T comparable]() *SinglyLinkedList[T] {
	newSinglyLinkedList := SinglyLinkedList[T]{head: nil, tail: nil}
	return &newSinglyLinkedList
}

func (s *SinglyLinkedList[T]) Add(val T) {
	if s.head == nil {
		var newNode *node[T] = NewNode(nil, val)
		s.head = newNode
		s.tail = newNode
	} else {
		var nextnode *node[T] = NewNode(nil, val)
		s.tail.next = nextnode
		s.tail = s.tail.next
	}

	length++
}

func (s *SinglyLinkedList[T]) Remove(value T) {
	if s.head == nil {
		return
	}
	if s.head.value == value {
		s.head = nil
		return
	}
	var prvNode *node[T] = s.head
	var currNode *node[T] = prvNode.next
	for currNode != nil {
		if length == 0 {
			break
		}
		if length == 1 {
			s.head = nil
		}
		if s.head.value == value {
			s.head = currNode
			break
		}
		if s.tail.value == value {
			prvNode.next = nil
			s.tail = prvNode
		}
		if currNode.value == value {
			prvNode.next = currNode.next
			currNode = nil
			break
		}
		prvNode = currNode
		currNode = currNode.next
	}
}

func (s *SinglyLinkedList[T]) PrintAll() {
	var headNode *node[T] = s.head
	for headNode != nil {
		fmt.Println(headNode.value)
		headNode = headNode.next
	}
}
