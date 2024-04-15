package main

func main() {
	// type assertion
	var i interface{} = "Hello, World"
	println(i)
	println(i.(string)) // afirmando que o tipo é string

	res, ok := i.(int) // testando se o tipo é int, ok vai retornar true ou false e res o valor
	println(res, ok) // se for falso, res vai receber zero
	
	res2 := i.(int) //se não receber o ok, o compilador vai retornar um panic
	println(res2)
}