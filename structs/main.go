package main

import "fmt"

// el simil a las clases
type Persona struct{
	Nombre string
	Edad uint8
	Email []string
}

func main() {
	// la forma tradicional de declarar
	var persona1 Persona
	persona1.Nombre = "Jose"
	persona1.Edad = 28

	fmt.Println(persona1.Nombre)

	// una forma mas sencilla de declarar
	persona2 := Persona{
		Nombre: "Luis",
		Edad: 15,
	}

	fmt.Println(persona2.Nombre, persona2.Edad)
	
	//haciendo una declaracion para luego llamarla en una estructura
	emailsDePersona3 := []string{"jyro777@gmail.com", "true.jhonnyr@gmail.com"}

	// toma el orden en el que fue declarado
	persona3 := Persona{
		"Pedro",
		17,
		emailsDePersona3,
	}

	fmt.Println(persona3.Nombre, persona3.Edad, persona3.Email[0])

}
