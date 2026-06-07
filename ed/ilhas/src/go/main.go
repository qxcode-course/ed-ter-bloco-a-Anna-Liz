package main

import (
	"bufio"
	"fmt"
	"os"
)

func rec(grid[][]byte, l, c int) {
	if l < 0 || c < 0 || l >= len(grid) || c >= len(grid[0]) || grid[l][c] == '0' {
		return
	}
	
	grid[l][c] = '0'

	rec(grid, l, c+1)
	rec(grid, l, c-1)
	rec(grid, l+1, c)
	rec(grid, l-1, c)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == '1' {
				count++
				rec(grid, i, j)
			}

		}
	}

	
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
