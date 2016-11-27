// Client (Cosumidor) del chat server que consuma los datos de lass gorutines y channels

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8005")
	if err != nil {
		log.Fatal(err)
	}

	// Definir variable que ignore errores
	done := make(chan struct{})

	// Definir una funcion que utiliza el paquete io
	go func() {
		io.Copy(os.Stdout, conn) // Ignorando errores
		log.Println("Terminamos")
		// Avisarle al gorutine principal
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()

	// Estamos esperando que la gorutine del background termine
	<-done
}

// Definir nueva funcion mustCopy
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
