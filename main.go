package main

import (
	"fmt"
)

//Hash table structure
const arraySize = 7

type HashTable struct {
	array [arraySize]*bucket
}

//bucket structure
type bucket struct {
	head *bucketNode
}

//bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

func (ht *HashTable) Insert(k string) {
	index := hash(k)
	ht.array[index].insert(k)
}

//Search
func (ht *HashTable) Search(k string) bool {
	index := hash(k)
	return ht.array[index].search(k)
}

//Delete
func (ht *HashTable) Delete(k string) {
	index := hash(k)
	ht.array[index].delete(k)
}

//insert
func (b *bucket) insert(k string) {

	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

//search - most return true if found
func (b *bucket) search(k string) bool {

	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
	}
}

//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

//Init
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result

}

func main() {
	fmt.Println("Hello - HashTable")

	nombres := [7]string{
		"XTBL",
		"NGEL",
		"LVIA",
		"TOBAL",
		"ALANA",
		"AGUELITA",
		"AGUELITO",
	}

	ht := Init()

	for _, v := range nombres {
		ht.Insert(v)
	}

	ht.delete("XTBL")
	ht.search("XTBL")
	ht.search("TOBAL")
	/*
		fmt.Println(ht)
		fmt.Println(hash("RANDY"))

		b := &bucket{}
		b.insert("Randy")
		b.insert("bamby")
		fmt.Println(b.search("bamby"))
		b.delete("bamby")
		fmt.Println(b.head.key)
		fmt.Println(b.search("brokoly"))
		fmt.Println(b.search("bamby"))
	*/
}
