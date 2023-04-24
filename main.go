package main

//This is minimal stack implementation inspired by
//https://youtu.be/fsbm1FOSDJ0
//if stack is empty it returns -1 on pop()
import (
	"fmt"
)

type Stack struct {
	items []int
}

//push adds a value at the end of the stack
func (s *Stack) push(element int) {
	s.items = append(s.items, element)

}

//pop removes a value at the end of the stack
func (s *Stack) pop() int {
	if len(s.items) > 0 {
		l := len(s.items) - 1
		result := s.items[l]
		s.items = s.items[:l]
		return result
	} else {
		return -1
	}
}

func main() {
	fmt.Println("Hola tio!!!")

	s := Stack{}
	s.push(1)
	s.push(2)
	s.push(9)
	s.push(10)
	fmt.Println("stack", s)
	for i := 0; i < 5; i++ {
		s.pop()
		fmt.Println("stack", s)
	}
}
