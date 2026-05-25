package main

import (
	"fmt"
	//"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	fmt.Print("[ ")
	for n := l.Front() ; n != l.End() ; n = n.Next() {
		if n == sword {
			fmt.Print(">", n.Value, " " )
		} else {
			fmt.Print(n.Value, " ")
		}
	}
	fmt.Println("]")
	return ""
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil || l.Size() == 0 {
		return nil
	}

	next := it.Next()

	if next == l.End() {
		next = l.Front()
	}

	return next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
