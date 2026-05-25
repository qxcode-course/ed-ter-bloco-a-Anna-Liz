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
	next *Node
	prev *Node
} 

type LList struct {
	head *Node
	tail *Node
	size int
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) SetValue(newValue int) {
	n.value = newValue
}
 
func (l *LList) Front() *Node {
	return l.head
}

func (l *LList) Back() *Node {
	return l.tail
}

func (l *LList) Size() int {
	return l.size
}

func NewLList() *LList {
	return &LList{}
}

func (l *LList) String() {
	fmt.Print("[")
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Next() == nil {
			fmt.Print(node.Value())
			break
		}
    	fmt.Print(node.Value(), ", ")
	}
	fmt.Println("]")
}

func (l *LList) Search(value int) *Node {
	for node := l.Front() ; node != nil ; node = node.Next() {
		if node.Value() == value {
			return node
		}
	}
	return nil
}

func (l *LList) Insert(node *Node, value int) {
	newNode := &Node{
		value: value,
		next: node,
		prev: node.prev,
	}
	if node.prev == nil {
		l.head = newNode
	} else {
		node.prev.next = newNode
	}
	node.prev = newNode
	l.size++
}

func (l *LList) Remove(node *Node) {
	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	l.size--
}

func (l *LList) PushBack(value int) {
	newNode := &Node{
		value: value,
		prev: l.tail,
	}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.size++
}

func (l *LList) PushFront(value int) {
	newNode := &Node{
		value: value,
		next: l.head,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		l.head = newNode
	}

	l.size++
}

func (l *LList) Walk() {
	for node := l.Front(); node != nil; node = node.Next() {
		fmt.Print(node.Value())
	} 

	for node := l.Back(); node != nil; node = node.Prev() {
		fmt.Printf("%v ", node.Value())
	}
}

func (l *LList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LList) PopBack() {
	if l.tail == nil {
		return
	}
	l.Remove(l.tail)
}

func (l *LList) PopFront() {
	if l.head == nil {
		return
	}
	l.Remove(l.head)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value())
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value())
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	node.value = newvalue
			} else {
			 	fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Remove(node)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
