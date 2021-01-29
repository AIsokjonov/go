package main
import "fmt"
import "unsafe"

type SinglyLinkedList struct {
	data float64 // this field can be any type
	next *SinglyLinkedList
}

type DoublyLinkedList struct {
	pr *DoublyLinkedList // predecessor
	data float64 // this field can be any type
	su *DoublyLinkedList // successor
}

type BinaryTree struct {
	left *BinaryTree
	data float64 // this field can be any type
	right *BinaryTree
}

func main() {
	fmt.Println("Size of Singly Linked List: ", unsafe.Sizeof(SinglyLinkedList{324.32, &SinglyLinkedList{}}))
	fmt.Printf("Size of Double Linked List: %v\n", unsafe.Sizeof(DoublyLinkedList{&DoublyLinkedList{}, 34.42, &DoublyLinkedList{}}))
	fmt.Printf("Size of Binary Tree: %v\n", unsafe.Sizeof(BinaryTree{&BinaryTree{}, 12.312, &BinaryTree{}}))
}
