package main

//A very simple Queue implementation. Based on https://youtu.be/fsbm1FOSDJ0

//---It is not production safe, There is not empty queue managment for Dequeue
import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {

	result := q.items[0]
	q.items = q.items[1:]

	return result

}

func main() {
	fmt.Println("Hola tío!!!")

	myQ := Queue{}
	fmt.Println("myQ", myQ)
	myQ.Enqueue(10)
	myQ.Enqueue(9)
	myQ.Enqueue(8)
	fmt.Println("myQ", myQ)
	fmt.Println("myQ", myQ)
	myQ.Dequeue()
	myQ.Dequeue()
	fmt.Println(myQ.Dequeue())
	fmt.Println("myQ", myQ)
}
