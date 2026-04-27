package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity > 0 {
		ms.capacity = ms.capacity * 2
	} else {
		ms.capacity = 1
	}
	newVector := make([]int, ms.capacity)

	for i := 0; i < ms.size; i++ {
		newVector[i] = ms.data[i]
	}

	ms.data = newVector
}

func (ms *MultiSet) binarySearch(value int) int {
	left := 0
	right := ms.size

	for left < right {
		mid := (left + right) / 2

		if value > ms.data[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	index := ms.binarySearch(value)

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) contains(value int) bool {
	index := ms.binarySearch(value)
	if ms.data[index] == value {
		return true
	} else {
		return false
	}
}

func (ms *MultiSet) erase(value int) {
	index := ms.binarySearch(value)
	if ms.data[index] == value {
		for i := index; i < ms.size-1; i++ {
			ms.data[i] = ms.data[i+1]
		}
		ms.size--
	} else {
		fmt.Println("value not found")
	}
}

func (ms *MultiSet) count(value int) int {
	if ms.size == 0 {
		return 0
	}
	index := ms.binarySearch(value)
	contador := 0
	for i := index; ms.data[i] == value; i++ {
		contador++
	}
	return contador
}

func (ms *MultiSet) unique() int {
	valorAtual := 0
	contador := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] != valorAtual {
			contador++
			valorAtual = ms.data[i]
		}
	}
	return contador
}

func (ms *MultiSet) clear() {
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println("[" + Join(ms.data[:ms.size], ", ") + "]")
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.count(value))
		case "unique":
			fmt.Println(ms.unique())
		case "clear":
			ms.clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
