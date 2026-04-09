package main

import (
	"fmt"
)

func aEhMaior(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func mdc(a, b int) int {
	var resto int

	if a == 0 || b == 0 {
		return a
	}

	if aEhMaior(a, b) {
		resto = a % b
		return mdc(b, resto)
	} else {
		resto = b % a
		return mdc(a, resto)
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
