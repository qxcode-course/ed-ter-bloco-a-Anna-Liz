package main

import "fmt"

func printar(vet []int) {
	if len(vet) == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < len(vet); i++ {
			if i == len(vet)-1 {
				fmt.Println(vet[i])
			} else {
				fmt.Print(vet[i], " ")
			}
		}
	}
}

func main() {
	var total, qtd int
	var possui []int
	fmt.Scan(&total, &qtd)

	for i := 0; i < qtd; i++ {
		var x int
		fmt.Scan(&x)
		possui = append(possui, x)
	}

	var repetidas []int
	for i := 1; i < qtd; i++ {
		if possui[i] == possui[i-1] {
			repetidas = append(repetidas, possui[i])
		}
	}

	var faltando []int
	var aux int
	for i := 1; i <= total; i++ {
		for aux < qtd && possui[aux] < i {
			aux++
		}
		if aux >= qtd || possui[aux] != i {
			faltando = append(faltando, i)
		}
	}

	printar(repetidas)
	printar(faltando)

}
