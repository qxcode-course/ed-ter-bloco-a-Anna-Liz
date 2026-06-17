package main
import "fmt"

func taNaLinha(matriz [][]int, linha int, valor int) bool {
    for i := 0; i < len(matriz); i++ {
        if matriz[linha][i] == valor {
            return true
        }
    }
    return false
}

func taNaColuna(matriz [][]int, coluna int, valor int) bool {
    for i := 0; i < len(matriz); i++ {
        if matriz[i][coluna] == valor {
            return true
        }
    }
    return false
}

func taNoQuadrado(matriz [][]int, linha, coluna, valor int) bool {
    n := len(matriz)

    tam := 3
    if n == 4 {
        tam = 2
    }

    inicioLin := (linha / tam) * tam
    inicioCol := (coluna / tam) * tam

    for i := inicioLin; i < inicioLin+tam; i++ {
        for j := inicioCol; j < inicioCol+tam; j++ {
            if matriz[i][j] == valor {
                return true
            }
        }
    }
    return false
}

func sudoku(matriz [][]int, id int) bool {
    n := len(matriz)

    if id == n*n {
        return true
    }

    linha := id / n
    coluna := id % n

    if matriz[linha][coluna] != 0 {
        return sudoku(matriz, id+1)
    }

    for i := 1; i <= n; i++ {
        if taNaColuna(matriz, coluna, i) {
            continue
        }

        if taNaLinha(matriz, linha, i) {
            continue
        }

        if taNoQuadrado(matriz, linha, coluna, i) {
            continue
        }

        matriz[linha][coluna] = i

        if sudoku(matriz, id+1) {
            return true
        }

        matriz[linha][coluna] = 0
    }
    
    return false
}

func main() {
    var n int
    fmt.Scan(&n)

    matriz := make([][]int, n)

    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)

        matriz[i] = make([]int, n)

        for j, valor := range linha {
            if valor == '.' {
                matriz[i][j] = 0
            } else {
                matriz[i][j] = int(valor - '0')
            }
        }
    }

    sudoku(matriz, 0)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Print(matriz[i][j])
        }
        fmt.Println()
    }
    
}