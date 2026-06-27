package main
import "fmt"

type Pos struct {
    l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l + 1, p.c}, 
		{p.l - 1, p.c}, 
		{p.l, p.c + 1}, 
		{p.l, p.c - 1},
	}
}

func inside(grid [][]int, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func main() {
    var m, n int
    fmt.Scan(&m, &n)

    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)

        for j := 0; j < n; j++ {
            fmt.Scan(&grid[i][j])
        }
    }

    var fila []Pos
    frescas := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i, j})
            } else {
                if grid[i][j] == 1 {
                    frescas++
                }
            }
        }
    }


    if frescas == 0 {
        fmt.Println(0)
        return
    }

    minutos := 0

    for len(fila) > 0 && frescas > 0 {
        tamanho := len(fila)

        for i := 0; i < tamanho; i++ {
            atual := fila[0]
            fila = fila[1:]
            for _, vizinhos :=  range atual.getNeig() {
                if inside(grid, vizinhos) && grid[vizinhos.l][vizinhos.c] == 1 {
                    grid[vizinhos.l][vizinhos.c] = 2
                    frescas--
                    fila = append(fila, vizinhos)
                }
            }
        }
        minutos++
    }

    if frescas > 0 {
        fmt.Println(-1)
    } else {
        fmt.Println(minutos)
    }
}