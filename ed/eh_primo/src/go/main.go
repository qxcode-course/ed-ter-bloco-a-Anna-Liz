package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	// aqui poderia ser x < 2
	if x == 1 {
		return false
	}
	// aqui poderia ser div*div, pois só precisamos verificar até a raiz quadrada
	// de x, tinha esquecido desse detalhe de discreta
	if div == x {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
