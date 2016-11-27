// Channels (Canales)

package main

import (
	"fmt"
	"sync"
)

const (
	gorutines = 20
	capacidad = 4
)

// Wg se usa para esperar que el programa termine.
var wg sync.WaitGroup

// recursos - buffered channel para manejar los strings

var recursos = make(chan string, capacidad)

// Main
func main() {

	// Agregar el numero de gorutines al waitgroup
	wg.Add(gorutines)

	// Lanzar las gorutines necesarias
	for gr := 1; gr <= gorutines; gr++ {
		go func(gr int) {
			worker(gr)
			wg.Done()
		}(gr)
	}

	// Añadimos los Strings

	for s := 'A'; s < 'A'+capacidad; s++ {
		recursos <- string(s)
	}

	wg.Wait()
}

// Lanzamos worker como una gorutine que procesa el trabajo del buffered channel

func worker(worker int) {
	// Recibir un string del channel
	valor := <-recursos

	fmt.Printf("Worker: %d: %s\n", worker, valor)

	recursos <- valor
}
