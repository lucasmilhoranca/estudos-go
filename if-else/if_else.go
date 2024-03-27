package main

import "fmt"

func main() {
	fmt.Println("If-else")
	number := 10

	if number > 5 {
		fmt.Println(number, "maior que 5")
	} else {
		fmt.Println(number, "menor que 5")
	}

	if x := number * number; x < 10 {
		fmt.Println(x, "menor que 10")
	} else {
		fmt.Println(x, "maior que 10")
	}

	var val interface{}
	val = "foo"
	if str, ok := val.(string); ok {

		fmt.Println(str, ok)
	}
}
