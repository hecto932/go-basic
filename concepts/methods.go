// Methods

package main

import "fmt"

// Jugador

type jugador struct {
	nombre   string
	goles    int
	partidos int
}

func (j *jugador) average() float64 {
	return float64(j.partidos) / float64(j.goles)
}

func main() {
	jugadores := []jugador{
		{"Carlos", 20, 60},
		{"Hector", 17, 30},
		{"Saurabh", 34, 60},
	}

	for _, jugador := range jugadores {
		fmt.Printf("Jugador: %s, Promedio: %f\n", jugador.nombre, jugador.average())
	}
}
