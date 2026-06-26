package main
import "fmt"

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

func nEsimo(n, atual, contador int) int {
    if primo(atual, 2) {
        contador++
        if contador == n {
            return atual
        }
    }

    return nEsimo(n, atual+1, contador)
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(nEsimo(n, 2, 0))
}