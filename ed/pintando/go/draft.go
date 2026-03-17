package main

import (
	"fmt"
	"math"
)

func main() {
	var ladoA, ladoB, ladoC float64
	fmt.Scan(&ladoA, &ladoB, &ladoC)

	var p = (ladoA + ladoB + ladoC) / 2
	var area = math.Sqrt(p * (p - ladoA) * (p - ladoB) * (p - ladoC))

	fmt.Printf("%.2f\n", area)
}
