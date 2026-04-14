package main
import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var fila []int
	for range n {
		var x int
		fmt.Scan(&x)
		fila = append(fila, x)
	}

	var m int
	fmt.Scanln(&m)
	remover := make([]int, 0, m)
	for range m {
		var y int
		fmt.Scan(&y)
		remover = append(remover, y)
	}
	fmt.Println(fila)
	fmt.Println(remover)
}
