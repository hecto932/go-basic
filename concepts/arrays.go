// Arrays

package main

import "fmt"

func main() {
	var nombres [5]string

	amigos := [5]string{"Raquel", "Isabel", "Fernando", "Enrique", "Jose"}

	nombres = amigos

	for i, nombre := range nombres {
		fmt.Println(nombre, &nombres[i])
	}
}
