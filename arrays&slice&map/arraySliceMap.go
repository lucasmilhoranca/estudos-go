package main

import "fmt"

func array() {
	fmt.Println("Arrays")

	var a [5]int //declara um array com 5 de comprimento. O comprimento do array faz parte do tipo!
	a[3] = 42
	i := a[3]
	fmt.Println(i)
	fmt.Println(a)

	var x = [2]int{1, 2}   //declara e inicializa
	y := [2]int{1, 2}      // forma abreviada
	z := [...]int{1, 2, 3} // ... -> compilador calcula o comprimento do array
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z, len(z))
}

func slice() {
	fmt.Println("Slices")

	var a []int                  // declara um slice - semelhante a um array, mas o comprimento não é especificado
	var b = []int{1, 2, 3, 4, 5} // declara e inicializa um slice (apoiado pelo array dado implicitamente)
	c := []int{1, 2, 3, 4, 5}    // forma abreviada
	chars := []string{0: "a", 1: "b", 2: "c"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(chars)

	a = append(a, 10, 20) // adiciona itens ao slice a
	f := append(a, b...)
	fmt.Println(a)
	fmt.Println(f)

	x := [3]string{"arg1", "arg2", "arg3"}
	s := x[:] //um slice referenciando o armazenamento do array x

	fmt.Println("Array", x)
	fmt.Println("Slice", s)
}

func mmap() {
	fmt.Println("Map")
	var m map[string]int
	m = make(map[string]int)
	m["key"] = 42
	fmt.Println(m["key"])

}

func main() {
	array()
	slice()
	mmap()
}
