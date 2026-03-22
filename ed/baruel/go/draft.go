package main

import "fmt"

func main() {
	var total, baruel int
	fmt.Scanln(&total)
	fmt.Scanln(&baruel)

	var figurinhas []int

	for i := 0; i < baruel; i++ {
		var x int
		fmt.Scan(&x)
		figurinhas = append(figurinhas, x)
	}

	var repetidas []int

	for j := 1; j < baruel; j++ {
		if figurinhas[j] == figurinhas[j-1] {
			repetidas = append(repetidas, figurinhas[j])
		}
	}

	var faltando []int
	aux := 0

	for h := 1; h <= total; h++ {
		for aux < baruel && figurinhas[aux] < h {
			aux++
		}

		if aux >= baruel || figurinhas[aux] != h {
			faltando = append(faltando, h)
		}
	}

	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		for v := range repetidas {
			if v == len(repetidas)-1 {
				fmt.Print(repetidas[v])
			} else {
				fmt.Print(repetidas[v], " ")
			}
		}
		fmt.Println()
	}

	if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		for v := range faltando {
			if v == len(faltando)-1 {
				fmt.Print(faltando[v])
			} else {
				fmt.Print(faltando[v], " ")
			}
		}
		fmt.Println()
	}
}
