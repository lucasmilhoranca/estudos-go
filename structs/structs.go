package main

import "fmt"

type Vertex struct {
	X, Y int
}

// Composição de structs
type Pessoa struct {
	name string
	age  int
	Enderaco
}

type Enderaco struct {
	city  string
	state string
}

//sem * troca apenas dentro da função
func (p *Pessoa) trocarNome(newName string) {
	p.name = newName
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

	pessoa := Pessoa{
		name: "Lucas",
		age: 20,
	}

	pessoa.Enderaco.city = "São Paulo"
	pessoa.state = "SP"
	fmt.Println(pessoa)

	pessoa.trocarNome("João")
	fmt.Println(pessoa)

	v.add(3)
	fmt.Println(v)
}