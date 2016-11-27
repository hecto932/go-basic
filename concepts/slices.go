// Slices

package main

import (
	"fmt"
)

func main() {
	var numeros []int

	for i := 0; i < 10; i++ {
		numeros = append(numeros, i*10)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	frutas := []string{"Manzana", "Naranja", "Pera", "Sandía", "Aguacate"}

	for i, fruta := range frutas {
		fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
	}

	slice := frutas[1:3]

	for i, fruta := range slice {
		fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
	}
}
