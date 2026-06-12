package main
import "fmt"


func backtracking(s []byte, pos int, n int) bool {
    if pos == len(s) {
        return true
    }
    if s[pos] != '.' {
        return backtracking(s, pos+1, n)
    }

    for i := 0; i <= n; i++ {
        ok := true

        inicio := pos - n
        if inicio < 0 {
            inicio = 0
        }

        fim := pos + n
        if fim >= len(s) {
            fim = len(s) - 1
        }

        for j := inicio; j <= fim; j++ {
            if j == pos {
                continue
            }
            if s[j] == byte('0'+ i) {
                ok = false
                break
            }
        }

        if ok {
            s[pos] = byte('0'+ i)

            if backtracking(s, pos+1, n) {
                return true
            }

            s[pos] = '.'
        }
    }
    return false
}

func main() {
    var sequencia string
    var n int

    fmt.Scan(&sequencia, &n)

    s := []byte(sequencia)
    backtracking(s, 0, n)

    fmt.Println(string(s))
    
}