package main

import "fmt"

//interface permite passar apenas assinatura de metodos, não permite atributos

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	name  string
	age   int
	ativo bool
	Enderaco
}

type Enderaco struct {
	city  string
	state string
}

func (c Cliente) Desativar() {
	c.ativo = false
	fmt.Printf("O cliente foi desativado: %s\n", c.name)
	fmt.Println(c.ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func showType(t interface{}) {
	fmt.Printf("Type: %T Valor: %v\n", t, t)
}

func main() {
	joao := Cliente{
		name:  "João",
		age:   20,
		ativo: true,
		Enderaco: Enderaco{
			city:  "São Paulo",
			state: "SP",
		},
	}
	fmt.Println(joao.ativo)
	Desativacao(joao)

	//interfaces vazias
	//simula tipagem dinamica

	var x interface{} = "hello"
	var y interface{} = 10

	showType(x)
	showType(y)

	x = 10
	y = "hello"
	showType(x)
	showType(y)
}
