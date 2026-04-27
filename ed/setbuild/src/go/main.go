package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func (v *Set) binarySearch(value int) int {
	left := 0
	right := v.size - 1

	for left <= right {
		mid := (left + right) / 2
		if v.data[mid] == value {
			return mid
		}
		if v.data[mid] > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func (v *Set) show() {
	if v.size == 0 {
		fmt.Println("[]")
	} else {
		fmt.Print("[")
		for i := range v.size {
			if i == v.size-1 {
				fmt.Print(v.data[i])
			} else {
				fmt.Print(v.data[i], ", ")
			}
		}
		fmt.Println("]")
	}

}

func (s *Set) grow() {
	newCapacity := 1
	if s.capacity > 0 {
		newCapacity = s.capacity * 2
	}

	newData := make([]int, newCapacity)
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}

	s.data = newData
	s.capacity = newCapacity
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Set) Insert(value int) {
	var index int

	if v.size == v.capacity {
		v.grow()
	}

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return
		}
		if v.data[i] < value {
			index++
		} else {
			break
		}
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++
}

func (v *Set) contains(value int) bool {
	if v.binarySearch(value) < 0 {
		return false
	} else {
		return true
	}
}

func (v *Set) erase(value int) {
	index := v.binarySearch(value)
	if index < 0 {
		fmt.Println("value not found")
		return
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]

	}

	v.size--
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			v.show()
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
