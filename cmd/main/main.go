package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}
type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}
func (d *doublyLinkedList) AddFrontNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
}
func (d *doublyLinkedList) AddEndNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
}
func (d *doublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}
func (d *doublyLinkedList) Size() int {
	return d.len
}
func (d *doublyLinkedList) ReverseDLL() {
	currentNode := d.head
	var nextInList *node
	d.head, d.tail = d.tail, d.head
	for currentNode != nil {
		nextInList = currentNode.next
		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next
		currentNode = nextInList
	}
}
func main() {
	doublyList := initDoublyList()
	fmt.Printf("Add Front Node: C\n")
	doublyList.AddFrontNodeDLL("C")
	fmt.Printf("Add Front Node: B\n")
	doublyList.AddFrontNodeDLL("B")
	fmt.Printf("Add Front Node: A\n")
	doublyList.AddFrontNodeDLL("A")
	fmt.Printf("Add End Node: D\n")
	doublyList.AddEndNodeDLL("D")
	fmt.Printf("Add End Node: E\n")
	doublyList.AddEndNodeDLL("E")
	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())
	err := doublyList.TraverseForward()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Reversing Doubly Linked List")
	doublyList.ReverseDLL()
	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())
	err = doublyList.TraverseForward()
	if err != nil {
		fmt.Println(err.Error())
	}
}
