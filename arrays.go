// Arrays

package main 

import "fmt"

func main() {
	var nombres [5]string 

	amigos:= [5]string{"Raquel", "Isabel", "Fernando", "Enrique", "Jose"}

	nombres = amigos

	fmt.Println(nombres)
	/*for i, nombres := range nombres {
		fmt.Println(nombres, &nombres[i])
	}*/
}