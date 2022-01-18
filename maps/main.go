package main

import "fmt"

func main() {
	// maps es para hacer un tipo de declarado de posicion personalizado (['es'] = "Español")
	/* idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Inglés"
	idiomas["jp"] = "Japonés"
	idiomas["fr"] = "Francés" */

	idiomas := map[string]string{
		"es": "Español",
		"en": "Inglés",
		"jp": "Japonés",
		"fr": "Francés",
	}
	if idioma, ok := idiomas["jp"]; ok {
		fmt.Println("la posicion 'jp' si exise y es", idioma)
	} else {
		fmt.Println("la posicion 'jp' no existe")
	}

	fmt.Println(idiomas)

	numeros := map[int]string{
		1: "uno",
		2: "dos",
		3: "tres",
		4: "cuatro",
		5: "cinco",
	}

	fmt.Println(numeros)
}
