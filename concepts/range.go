// Range

package main

import "fmt"

func main() {
	numeros := []int{2, 4, 6}
	suma := 0
	for _, num := range numeros {
		suma += num
	}

	fmt.Println("Suma:", suma)

	for i, numero := range numeros {
		if numero == 3 {
			fmt.Println("Index: ", i)
		}
	}

	algo := map[string]string{"a": "auto", "b": "bebe"}

	for key, value := range algo {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
