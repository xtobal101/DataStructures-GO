package main

//A very simple implementation of a linked list data structure"

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) preppend(new *node) {
	second := l.head
	l.head = new
	new.next = second
	l.length++
}

func (l linkedList) printListData() { // The reciever is gets a copy of the value not of the adress. It is not a pointer

	currNode := l.head

	for l.length != 0 {
		fmt.Printf("%d ", currNode.data)
		currNode = currNode.next
		l.length-- //It is possible to modify this value without affecting the original list becouse it is a copy
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteByValue(v int) {

	if l.length == 0 {
		return
	}

	if l.head.data == v {
		l.head = l.head.next
		l.length--
	} else {

		prevToDelete := l.head

		for prevToDelete.next.data != v {
			if prevToDelete.next.next == nil {
				return
			}
			prevToDelete = prevToDelete.next
		}

		prevToDelete.next = prevToDelete.next.next
		l.length--
	}
}

func main() {
	fmt.Println("Hola tio")

	ll := &linkedList{}

	n1 := &node{data: 84}
	n2 := &node{data: 42}
	n3 := &node{data: 21}

	n11 := &node{data: 8484}
	n22 := &node{data: 4242}
	n33 := &node{data: 2121}

	ll.preppend(n1)
	ll.preppend(n2)
	ll.preppend(n3)

	ll.preppend(n11)
	ll.preppend(n22)
	ll.preppend(n33)

	ll.printListData()
	ll.deleteByValue(2)
	ll.deleteByValue(100)
	ll.deleteByValue(8484)
	ll.printListData()

}
