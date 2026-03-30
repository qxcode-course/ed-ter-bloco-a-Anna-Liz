package main

import "fmt"

func printFila(fila []int, pos int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		if i == pos {
			fmt.Printf("%d> ", fila[i])
		} else {
			fmt.Printf("%d ", fila[i])
		}
	}
	fmt.Println("]")
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	var fila []int

	for i := 1; i <= n; i++ {
		fila = append(fila, i)
	}

	pos := e - 1
	for len(fila) > 1 {
		printFila(fila, pos)

		morte := (pos + 1) % len(fila)

		fila = append(fila[:morte], fila[morte+1:]...)

		pos = morte % len(fila)
	}

	printFila(fila, pos)

}
