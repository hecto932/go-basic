// Chat server
// Trabajo pesado

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Desarrollando el componente que emite los mensajes
// Los mensajes salientes

type client chan<- string

// Los mensajes entrantes
var (
	entrantes = make(chan client)
	salientes = make(chan client)
	mensajes  = make(chan string)
)

// 2- funcion broadcaster

// Todos los clientes conectados
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-mensajes:
			// Emitir el mensaje entrane a todos los clientes que estan cconectados de channels
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entrantes:
			clients[cli] = true

		case cli := <-salientes:
			delete(clients, cli)
			close(cli)
		}
	}
}

// 3.- Handler o manejador de datos.

func handleConn(conn net.Conn) {
	// Mensajes salientes
	ch := make(chan string)
	go clientWriter(conn, ch)

	quien := conn.RemoteAddr().String()
	ch <- "Tú eres" + quien

	mensajes <- quien + "se ha conectado a la sala"
	entrantes <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		mensajes <- quien + ":" + input.Text()
	}

	salientes <- ch
	mensajes <- quien + "se ha ido"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8005")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept() // Recibimos el mensaje
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
