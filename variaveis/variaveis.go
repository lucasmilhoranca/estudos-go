package main

import "fmt"

/*
	Tipos integrados
	bool
	string
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte
	rune
	float32 float64
	complex64 complex128
*/

func main() {
	fmt.Println("Variáveis")

	// bool
	var booleano bool = true
	fmt.Println(booleano)

	// string
	var texto string = "Olá, mundo!"
	fmt.Println(texto)

	// int
	var inteiro int = 10
	fmt.Println(inteiro)

	// int8
	var inteiro8 int8 = 127
	fmt.Println(inteiro8)

	// int16
	var inteiro16 int16 = 32767
	fmt.Println(inteiro16)

	// int32
	var inteiro32 int32 = 2147483647
	fmt.Println(inteiro32)

	// int64
	var inteiro64 int64 = 9223372036854775807
	fmt.Println(inteiro64)

	// uint
	var uinteiro uint = 10
	fmt.Println(uinteiro)

	// uint8
	var uinteiro8 uint8 = 255
	fmt.Println(uinteiro8)

	// uint16
	var uinteiro16 uint16 = 65535
	fmt.Println(uinteiro16)

	// uint32
	var uinteiro32 uint32 = 4294967295
	fmt.Println(uinteiro32)

	// uint64
	var uinteiro64 uint64 = 18446744073709551615
	fmt.Println(uinteiro64)

	// uintptr
	var ponteiro uintptr = 0xFFFF
	fmt.Println(ponteiro)

	// byte (alias para uint8)
	var byteVar byte = 255
	fmt.Println(byteVar)

	// rune (alias para int32)
	var runeVar rune = 'A'
	fmt.Println(runeVar)

	// float32
	var flutuante32 float32 = 3.14
	fmt.Println(flutuante32)

	// float64
	var flutuante64 float64 = 3.141592653589793
	fmt.Println(flutuante64)

	// complex64
	var complexo64 complex64 = complex(5, 2)
	fmt.Println(complexo64)

	// complex128
	var complexo128 complex128 = complex(10, 3)
	fmt.Println(complexo128)
}