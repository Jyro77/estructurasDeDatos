package main

import "fmt"

func main() {

	edades := [5]int8{
		18,
		25,
		30,
		23,
		29,
	}

	// el valor de arrays vacios para int, string y boolean

	var bol [3]bool
	var num [3]int8
	var str [3]string
	size := len(edades)

	fmt.Println(bol, num, str)
	fmt.Println("El tamaño de edades es:", size)
	fmt.Println(edades)
}
