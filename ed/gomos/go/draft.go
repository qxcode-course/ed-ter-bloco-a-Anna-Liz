package main

import "fmt"

type Par struct {
	x int
	y int
}

func main() {
	var qtd int
	var direcao string
	fmt.Scan(&qtd, &direcao)

	var cobrinha []Par
	for i := 0; i < qtd; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		cobrinha = append(cobrinha, Par{x, y})
	}

	var deslocX, deslocY int
	switch direcao {
	case "L":
		deslocX = -1
		deslocY = 0
	case "R":
		deslocX = 1
		deslocY = 0
	case "U":
		deslocX = 0
		deslocY = -1
	case "D":
		deslocX = 0
		deslocY = 1
	}

	novoPar := Par{
		x: cobrinha[0].x + deslocX,
		y: cobrinha[0].y + deslocY,
	}
	cobrinha = append([]Par{novoPar}, cobrinha...)
	cobrinha = cobrinha[:len(cobrinha)-1]

	for _, j := range cobrinha {
		fmt.Println(j.x, j.y)
	}
}
