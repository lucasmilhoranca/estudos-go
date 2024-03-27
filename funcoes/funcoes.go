package main

import (
	"fmt"
)

func funcaoSimples() {
	fmt.Println("Função simples")
}

func funcaoParametrizada(nome string) {
	fmt.Println("Olá,", nome)
}

func funcaoParametrizadaMesmoTipo(nome, sobrenome string) {
	fmt.Println("Olá,", nome, sobrenome)
}

func funcaoRetorno() string {
	return "Função com retorno"
}

func funcaoRetornoMultiplosValores() (string, string) {
	return "Função com retorno", "Multiplos valores"
}

func funcaoRetornoNomeado() (primeiroRetorno string, segundoRetorno string) {
	primeiroRetorno = "Retorno 1"
	segundoRetorno = "Retorno 2"
	return
}

func funcaoVariadica(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func scope() func() int {
	outerVar := 2
	foo := func() int {
		return outerVar
	}
	return foo
}

func outer() (func() int, int) {
	outerVar := 2
	inner := func() int {
		outerVar += 99
		return outerVar
	}
	inner()
	return inner, outerVar
}

func erro() (int, error) { // Não há tratamento de exceção. Funções que podem produzir um erro apenas declaram um valor de retorno adicional do tipo Error
	return 0, nil
}

func main() {
	//atribuindo uma função a uma variável estilo => no javascript
	funcao := func(a, b int) int {
		return a + b
	}
	fmt.Println(funcao(1, 2))

	funcaoSimples()

	funcaoParametrizada("Jhon")

	funcaoParametrizadaMesmoTipo("Jhon", "Doe")

	fmt.Println(funcaoRetorno())

	fmt.Println(funcaoRetornoMultiplosValores())

	fmt.Println(funcaoRetornoNomeado())

	fmt.Println(funcaoVariadica(1, 2, 3, 4, 5))

	foo := scope()
	fmt.Println(foo())

	inner, outer := outer()
	fmt.Println(inner())
	fmt.Println(outer)

	result, err := erro()
	fmt.Println(result)
	fmt.Println(err)
}
