package main

import "data-structure-algorithm/singlylinkedlist"

func main() {
	TestSinglyLinkedList()
}

// SinglyLinkedList
func TestSinglyLinkedList() {
	var nums *singlylinkedlist.SinglyLinkedList[int] = singlylinkedlist.NewSinglyLinkListed[int]()
	nums.Add(1)
	nums.Remove(1)
	nums.PrintAll()
}
