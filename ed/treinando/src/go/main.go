package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostrAux(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	x := strconv.Itoa(vet[0])
	if len(vet) > 1 {
		vet = vet[1:]
		return x + ", " + tostrAux(vet)
	} else {
		return x
	}
}

func tostr(vet []int) string {
	return "[" + tostrAux(vet) + "]"
}

func tostrrevAux(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	x := strconv.Itoa(vet[len(vet)-1])
	if len(vet) > 1 {
		vet = vet[:len(vet)-1]
		return x + ", " + tostrrevAux(vet)
	} else {
		return x
	}
}

func tostrrev(vet []int) string {
	return "[" + tostrrevAux(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) > 1 {
		vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
		vet = vet[:len(vet)-1]
		vet = vet[1:]
		reverse(vet)
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}

	x := vet[0]
	vet = vet[1:]

	return x + sum(vet)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	x := vet[0]
	vet = vet[1:]

	return x * mult(vet)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice

func min(vet []int) int {
	var rec func(v []int) (int, int)

	if len(vet) == 0 {
		return -1
	}

	rec = func(v []int) (int, int) {

		if len(v) == 1 {
			return v[0], 0
		}

		val, idx := rec(v[1:])

		if v[0] <= val {
			return v[0], 0
		}

		return val, idx + 1
	}
	_, idx := rec(vet)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
