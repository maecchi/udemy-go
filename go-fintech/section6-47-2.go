package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func question2() {
	v := Vertex{3, 4}
	fmt.Println(v)
}

func main() {
	question2()
}