// Mapas

package main

import (
	"fmt"
)

func main() {

	departamentos := make(map[string]int)

	departamentos["Devs"] = 25
	departamentos["Marketing"] = 50
	departamentos["Ejecutivos"] = 4
	departamentos["Ventas"] = 60
	departamentos["Mantenimiento"] = 8

	for key, value := range departamentos {
		fmt.Printf("Departamento: %s, Personas: %d\n", key, value)
	}
}
