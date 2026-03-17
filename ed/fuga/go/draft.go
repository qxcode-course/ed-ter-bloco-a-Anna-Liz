package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)

	for F != P {
		F += D
		if F > 15 {
			F = 0
		}

		if F < 0 {
			F = 15
		}

		if F == H {
			fmt.Println("S")
			return
		}
	}

	fmt.Println("N")
}
