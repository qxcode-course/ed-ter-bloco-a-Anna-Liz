package main

import (
	"fmt"
)

func primo(n, div int) bool {
    if n < 2 {
        return false
    }
    if div*div > n {
        return true
    }
    if n % div == 0 {
        return false
    }
    return primo(n, div+1)
}

func nPrimos(vetor []int, atual, pos int) {
    if pos == len(vetor) {
        return
    }

    if primo(atual, 2) {
        vetor[pos] = atual
        nPrimos(vetor, atual+1, pos+1)
    } else {
        nPrimos(vetor, atual+1, pos)
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]int, n)
    nPrimos(vetor, 2, 0)
    fmt.Print("[")
    for i := 0; i < len(vetor); i++ {
        if i == len(vetor)-1{
            fmt.Print(vetor[i])
        } else {
            fmt.Print(vetor[i], ", ") 
        }
    }
    fmt.Println("]")
}