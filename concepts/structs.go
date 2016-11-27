// Structs

package main

import (
	"fmt"
)

type usuario struct {
	nombre    string
	direccion string
	edad      int
}

func main() {

	hector := usuario{
		nombre:    "Hector",
		direccion: "Mexico Cty, Mexico",
		edad:      24,
	}

	fmt.Println("Nombre", hector.nombre)
	fmt.Println("Direccion", hector.direccion)
	fmt.Println("Edad", hector.edad)

	nicole := struct {
		nombre    string
		direccion string
		edad      int
	}{
		nombre:    "Nichole",
		direccion: "Calle 13",
		edad:      22,
	}

	fmt.Println("Nombre", nicole.nombre)
	fmt.Println("Direccion", nicole.direccion)
	fmt.Println("Edad", nicole.edad)
}
