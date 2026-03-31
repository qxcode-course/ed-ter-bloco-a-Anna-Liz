package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absolute(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

/*
	Pesquisando descobri que fazendo com append não é muito eficiente,

pois é preciso realocar a memória várias vezes para acrescentar os
outros elementos.

Outra forma mais eficiente poderia ser:
  - criar um slice de lista_homens do tamanho do vetor inicial, afinal
    sabemos que ela nao poderá ser maior, e dessa forma garantimos que
    haverá realocação
    :D
*/
func getMen(vet []int) []int {
	var lista_homens []int
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			lista_homens = append(lista_homens, vet[i])
		}
	}
	return lista_homens
}

// Podemos fazer make([]int, 0, len(vet)) para criar um slice com 0 elementos
func getCalmWomen(vet []int) []int {
	lista_mulheres := make([]int, 0, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && -(vet[i]) < 10 {
			lista_mulheres = append(lista_mulheres, vet[i])
		}
	}
	return lista_mulheres
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet)-1; j++ {
			if vet[j] > vet[j+1] {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}

	return vet
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet)-1; j++ {
			if absolute(vet[j]) > absolute(vet[j+1]) {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func reverse(vet []int) []int {
	lista := make([]int, len(vet))
	copy(lista, vet)

	var aux = len(lista) - 1
	for i := 0; i < len(lista)/2; i++ {
		lista[i], lista[aux] = lista[aux], lista[i]
		aux--
	}
	return lista
}

func unique(vet []int) []int {
	_ = vet
	return nil
}

func repeated(vet []int) []int {
	_ = vet
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
