package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Abs() int {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vertex) add(n int) {
	v.X += n
	v.Y += n
}

func main() {
	var v = Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	fmt.Println(v.Abs())

	v.add(3)
	fmt.Println(v)
}