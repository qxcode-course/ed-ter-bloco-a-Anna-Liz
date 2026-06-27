package main
import "fmt"

func degraus(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    return degraus(n-1) + degraus(n-3)
}
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(degraus(n))
}