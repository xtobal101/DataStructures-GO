package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) Print() {

	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)

		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("hola tio!!!")
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddVertex(0)

}
