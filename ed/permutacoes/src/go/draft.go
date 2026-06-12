package main
import "fmt"

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
    permutacao("", palavra)
    
}