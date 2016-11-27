// GoRutines

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Se ejecuta antes de Main
func init() {
	// Alojando/reservando (allocate) un procesador logico para que lo hice el scheduler.
	runtime.GOMAXPROCS(1)
}

// Main - Entrada a los programas
func main() {
	// Wait group: Es un grupo de espera para iniciar el conteo de las dos(2) gorutinas
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Iniciar Gorutines")

	// Creando funcion anonima y una gorutine

	go func() {
		for i := 100; i >= 0; i-- {
			fmt.Printf("[A: %d]\n", i)
		}
		//Avisar a main la culminacion
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("[B: %d]\n", i)
		}
		//Avisar a main la culminacion
		wg.Done()
	}()

	//Esperar a que termine las gorutines
	fmt.Println("Esperando a que termine las gorutines...")

	wg.Wait()

	fmt.Println("\nPrograma finalizado.")

}
