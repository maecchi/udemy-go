package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func question1() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}

func main() {
	question1()
}