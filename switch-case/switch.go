package main

import "fmt"

func main() {
	fmt.Println("Switch Case")

	switch number := 2; number {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}