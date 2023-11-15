package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func insert(head *Node, data int) *Node {
	n := &Node{Data: data}
	if head == nil {
		return n
	} else {
		n.Next = head
		return n
	}
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.Data, " -> ")
		head = head.Next
	}
	fmt.Println(nil)
}

func main() {
	var link *Node
	link = insert(link, 1)
	link = insert(link, 2)
	link = insert(link, 3)
	link = insert(link, 4)

	PrintList(link)
}
