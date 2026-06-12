package main
import "fmt"

func backtracking(elem []int, alvo int, soma int, id int) bool{
    if soma == alvo{
        return true
    }
    if soma > alvo{
        return false
    }
    if id == len(elem){
        return false
    }
    if backtracking(elem, alvo, soma+elem[id], id+1) {
        return true
    }
    if backtracking(elem, alvo, soma, id+1) {
        return true
    }
    return false
}

func main() {
    var qtd, alvo int
    fmt.Scan(&qtd, &alvo)

    elem := make([]int, qtd)

    for i := 0; i < qtd; i++ {
        var valor int
        fmt.Scan(&valor)
        elem[i] = valor
    }

    if (backtracking(elem, alvo, 0, 0)) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}