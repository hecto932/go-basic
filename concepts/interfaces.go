// Interfaces

package main

import "fmt"

type speaker interface {
	speak()
}

type ingles struct{}

func (ingles) speak() {
	fmt.Println("Hello world!")
}

type chino struct{}

func (chino) speak() {
	fmt.Println("你好世界")
}

func main() {
	var sp speaker

	var i ingles
	sp = i
	sp.speak()

	var c chino
	sp = c
	sp.speak()

	decirHola(new(ingles))
	decirHola(&chino{})
}

func decirHola(sp speaker) {
	sp.speak()
}
