package main
import "fmt"
import "sort"

func permutacao(primeira string, resto string) {
    if len(resto) == 0 {
        fmt.Println(primeira)
        return
    }
    for i := 0; i < len(resto); i++ {
        permutacao(primeira+string(resto[i]), resto[:i]+resto[i+1:])
    }
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    r := []rune(palavra) 
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    
    permutacao("", string(r))
    
}