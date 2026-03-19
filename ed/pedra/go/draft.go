package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var ganhador int
	var aux int = 100

	for i := range n {
		var A, B int
		fmt.Scanln(&A, &B)
		var diferenca = A - B

		if A > 9 && B > 9 {
			if diferenca < 0 {
				diferenca = -diferenca
			}

			if diferenca < aux {
				aux = diferenca
				ganhador = i
			}
		}
	}

	if ganhador == 0 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(ganhador)
	}
}
