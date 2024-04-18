package main

import "fmt"

func main() {
	fmt.Println("Loops")
	//linguagem go não tem while ou do-while, apenas for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}