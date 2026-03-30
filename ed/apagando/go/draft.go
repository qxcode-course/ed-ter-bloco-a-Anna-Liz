package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)

	var fila []int
	for range n {
		var x int
		fmt.Scan(&x)
		fila = append(fila, x)
	}

	fmt.Scan(&m)
	desistentes := make(map[int]struct{})
	for range m {
		var x int
		fmt.Scan(&x)
		desistentes[x] = struct{}{}
	}

	var fila_final []int
	for i := 0; i < n; i++ {
		_, existe := desistentes[fila[i]]
		if existe == false {
			fila_final = append(fila_final, fila[i])
		}

	}

	for j := 0; j < len(fila_final); j++ {
		fmt.Print(fila_final[j], " ")
	}
	fmt.Println()
}
