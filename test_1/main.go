package main

import (
	"fmt"
)

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range lista {
		if lista[i]%2 == 0 {
			fmt.Println(lista[i], "is even")
		}
		if lista[i]%2 != 0 {
			fmt.Println(lista[i], "is odd")
		}
	}
}
