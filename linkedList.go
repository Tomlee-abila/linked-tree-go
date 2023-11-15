package main

import "fmt"

type Node struct {
	Data int
	Data2 string
	Next *Node
}

func insert(head *Node, data int, data2 string) *Node {
	n := &Node{Data: data, Data2: data2}
	if head == nil {
		return n
	} else {
		n.Next = head
		return n
	}
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.Data, " -> ", head.Data2," -> ")
		head = head.Next
	}
	fmt.Println(nil)
}

func main() {
	var link *Node
	link = insert(link, 1, "Hello")
	link = insert(link, 2, "how")
	link = insert(link, 3, "are")
	link = insert(link, 4, "you")

	PrintList(link)
}
