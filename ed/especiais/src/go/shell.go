package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	contador := make(map[int]int)

	for _, x := range vet {
		stress := abs(x)
		contador[stress]++
	}

	var par []Pair
	for s := 1; s <= 100; s++ {
		valor, existe := contador[s]
		if existe {
			par = append(par, Pair{s, valor})
		}
	}

	return par
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}
	
	var par []Pair

	atual := vet[0]
	qtd := 1
	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			qtd++
		} else {
			par = append(par, Pair{atual, qtd})
			qtd = 1
			atual = vet[i]
		}
	}
	par = append(par, Pair{atual, qtd})
	return par
}

func mnext(vet []int) []int {
	mapa := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if i > 0 && vet[i-1] < 0 {
				mapa[i] = 1
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				mapa[i] = 1
			}
		}
	}
	return mapa
}

func alone(vet []int) []int {
	mapa := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			continue
		}

		left := i > 0 && vet[i-1] < 0
		right := i < len(vet)-1 && vet[i+1] < 0

		if !left && !right {
			mapa[i] = 1
		}
	}

	return mapa
}

func couple(vet []int) int {
	homens := make(map[int] int)
	mulheres := make(map[int] int)

	for _, x := range vet {
		if x > 0 {
			homens[x]++
		} else {
			mulheres[-x]++
		}
	}

	total := 0

	for id, qtdH := range homens {
		qtdM := mulheres[id]

		if qtdH < qtdM {
			total = total + qtdH
		} else {
			total = total + qtdM
		}
		
	}
	return total
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	var resultante []int

	for i := 0; i < len(vet); i++ {
		mantem := true
		for j := 0; j < len(posList); j++ {
			if i == posList[j] {
				mantem = false
				continue
			}
		}

		if mantem {
			resultante = append(resultante, vet[i])
		}
	}
	return resultante
}

func clear(vet []int, value int) []int {
	var resultante []int
	for _, x := range vet {
		if x != value {
			resultante = append(resultante, x)
		}
	}
	return resultante
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
