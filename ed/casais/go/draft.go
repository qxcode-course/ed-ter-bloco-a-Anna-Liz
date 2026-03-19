package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var descasados []int
	var contagem int

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		descasados = append(descasados, x)

		for j := range descasados {
			if descasados[j] == -x {
				descasados[i] = 0
				descasados[j] = 0
				contagem++
			}
		}
	}

	fmt.Println(contagem)
}
