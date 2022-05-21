package main

import (
	"fmt"
)

// if you can make with golang language, try another language
func mainBfs() {
	var q,m,n int
	fmt.Scanf("%d", &q)
	for ;q != 0;q--{
		fmt.Scanf("%m %n", &m, &n)
		// res := make(map[int32])
		for ;m != 0;m--{

		}
	}

	TestGraphBfs()
}

func TestGraphBfs(){
	edges := [][]int32{
		1:{2,3},
		2:{1},
		3:{4,5},
		4:{2,3},
		5:{},
	}
	fmt.Println(edges[1])
}

func TestGraph() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)

	test.Print()
}

// Print adjacent list
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println("")
}

//Graphs
type Graph struct {
	// array pointer vertex
	vertices []*Vertex
}

//Vertex adjacency list
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//AddEdge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//AddVertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//check wether vertex has contain key
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}



func Bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// Build graph
	graph := &Graph{}
	for _,v := range edges{
		v0, v1 := int(v[0]), int(v[1])
		graph.AddVertex(v0)
		graph.AddEdge(v0,v1)

		graph.AddVertex(v1)
		graph.AddEdge(v1,v0)
	}

	graph.Print()

	return []int32{0}
}
