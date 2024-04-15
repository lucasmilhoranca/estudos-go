package main

func main() {
	// Memórioa -> Endereço -> Valor

	a := 10
	println(&a, a)
	var ponteiro *int = &a
	println(ponteiro, *ponteiro)
	*ponteiro = 20
	println(a, ponteiro, *ponteiro)
	b := &a
	println(b, *b)
}
