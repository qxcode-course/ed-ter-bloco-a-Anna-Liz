package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, word string, l, c, id int) bool {
	qtdL := len(grid)
	qtdC := len(grid[0]) 

	if id == len(word) {
		return true
	}

	if l < 0 || l >= qtdL || c < 0 || c >= qtdC {
		return false
	}

	if grid[l][c] != word[id] {
		return false
	}

	char := grid[l][c]
	grid[l][c] = '.'

	vizinhos := 
		dfs(grid, word, l+1, c, id+1) ||
		dfs(grid, word, l-1, c, id+1) ||
		dfs(grid, word, l, c+1, id+1) ||
		dfs(grid, word, l, c-1, id+1)

	grid[l][c] = char
	
	return vizinhos
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[0]); c++ {
			if dfs(grid, word, l, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
