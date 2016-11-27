// Funciones

package main

import (
	"fmt"
)

type usuario struct {
	nombre string
	email  string
}

func nuevoUsuario() (*usuario, error) {
	return &usuario{"Hector", "hecto932@gmail.com"}, nil
}

func main() {
	u, err := nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*u)

	_, err = nuevoUsuario()

	if err != nil {
		fmt.Println(err)
		return
	}

}
