package main

import "fmt"

func main() {

	//practica de Slice
	//var nombres []string
	// make() se le pasan 3 parametros, 1ro tipo del slice, 2do tamaño inicial del slice, 3ro la capacidad inicial del slice * este último es opcional.
	nombres := make([]string, 0) // este es el metodo mas usado

	/* 	nombres := []string { // este es otra opcion para crear los slice
		"Luicito",
		"Pablito",
		"Jorge",
	} */

	nombres = append(nombres, "Marcos")
	nombres = append(nombres, "Joshua")
	nombres = append(nombres, "Oscar")
	nombres = append(nombres, "Mateo")
	nombres = append(nombres, "Lucas")
	nombres = append(nombres, "Jose")
	nombres = append(nombres, "Pedro")
	nombres = append(nombres, "Luis")
	nombres = append(nombres, "Miguel")
	fmt.Println("algo", nombres)
	nombres[5] = "Roberto"

	fmt.Printf("el tamaño es de %d, y la capacidad es de %d\n", len(nombres), cap(nombres))
	fmt.Println("algo", nombres)
}
