package main
import "fmt"
func main() {
    var entrada string
    fmt.Scan(&entrada)

    texto := []rune{}
    cursor := 0

    for _, char := range entrada {
        switch char {
        case 'R':
            texto = append(texto, '\n')
            cursor++
        case 'B':
            if cursor != 0 {
                texto = append(texto[:cursor-1], texto[cursor:]...)
                cursor--
            }
        case 'D':
            if cursor < len(texto) {
                texto = append(texto[:cursor], texto[cursor+1:]...)
            }
        case '>':
            if cursor < len(texto) {
                cursor ++
            }
        case '<':
            if cursor != 0 {
                cursor --
            }
        default: 
            texto = append(texto[:cursor], append([]rune{char}, texto[cursor:]...)...)
            cursor++
        }

    }

    for i := 0; i < len(texto); i++ {
        if i == cursor  {
            fmt.Print("|")
        }

        if i < len(texto) {
            fmt.Printf("%c", texto[i])
        }
    }
    if cursor == len(texto) {
        fmt.Print("|")
    }

    fmt.Println()

}
