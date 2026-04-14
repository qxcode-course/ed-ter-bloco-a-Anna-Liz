package main

import (
	"fmt"
)

func main() {
	var n, contador int
	var descasados []int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		descasados = append(descasados, x)

		for j := 0; j < len(descasados); j++ {
			if descasados[j] == -x {
				descasados[i] = 0
				descasados[j] = 0
				contador++
			}
		}
	}

	fmt.Println(contador)

}
