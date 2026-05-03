package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DList struct {
	head *Node
	size int
}

func NewDList() *DList {
	dlist := &DList{}
	dlist.head = &Node{}
	dlist.size = 0
	dlist.head.next = dlist.head
	dlist.head.prev = dlist.head
	return dlist
}

func (list *DList) String() {
	iterador := list.head.next
	fmt.Print("[")
	for {
		if iterador == list.head {
			break
		}

		if iterador == list.head.prev {
			fmt.Print(iterador.value)
			break
		}

		fmt.Print(iterador.value, ", ")
		iterador = iterador.next
	}
	fmt.Println("]")
}

func (list *DList) Insert(B *Node, value int) {
	A := B.prev
	C := &Node {
		value: value,
		next: B,
		prev: A,
	} 
	A.next = C
	B.prev = C
	list.size++
}

func (list *DList) Remove(B *Node) {
	if B == list.head{
		return
	}
	A := B.prev
	C := B.next
	A.next = C
	C.prev = A

	B.next = nil
	B.prev = nil
}

func (list *DList) PushFront(value int) {
	list.Insert(list.head.next, value)
}

func (list *DList) PushBack(value int) {
	list.Insert(list.head, value)
}

func (list *DList) PopFront(){
	list.Remove(list.head.next)
}

func (list *DList) PopBack(){
	list.Remove(list.head.prev)
}

func (list *DList) Size() int{
	return list.size
}

func (list *DList) Clear() {
	list.size = 0
	list.head.next = list.head
	list.head.prev = list.head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewDList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			ll.String()
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
