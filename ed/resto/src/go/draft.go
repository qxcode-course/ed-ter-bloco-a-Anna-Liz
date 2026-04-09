package main

import "fmt"

func resto(numero int) {
	if numero == 0 {
		return
	}
	resti := numero % 2
	numero = numero / 2
	resto(numero)
	fmt.Println(numero, resti)
}

func main() {
	var x int = 0
	fmt.Scan(&x)
	resto(x)
}
