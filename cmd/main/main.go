package main

import (
	"fmt"
)

type Node struct {
	Previous *Node
	Data     int
	Next     *Node
}

var tail *Node

func (n *Node) AddNode(data int) {
	fmt.Println("Inserting node", data)
	if tail == nil {
		tail = n
	}
	n = tail
	newNode := Node{n, data, nil}
	n.Next = &newNode
	tail = &newNode
}

func (n *Node) PrintNode() {
	fmt.Println("Printing nodes")
	iter := n
	for iter != nil {
		fmt.Println(iter.Data)
		iter = iter.Next
	}
}

func (n *Node) DeleteLast() {
	fmt.Println("Deleting last node", tail.Data)
	tail = tail.Previous
	(tail.Next).Previous = nil
	tail.Next = nil
}

func main() {
	newNode := Node{nil, 10, nil}
	newNode.AddNode(20)
	newNode.AddNode(30)
	newNode.AddNode(40)
	newNode.PrintNode()
	newNode.DeleteLast()
	newNode.PrintNode()
	newNode.AddNode(50)
	newNode.AddNode(60)
	newNode.AddNode(70)
	newNode.PrintNode()
	newNode.DeleteLast()
	newNode.PrintNode()
}
