package main

import (
	"fmt"
	"./linkedlist"
	"./fibonacci"
)

func main() {
	FibonacciSequence()
}

func LinkedList() {
	values := []int64{10, 235, 2975, 37234, 108439}
	head := ll.CreateList(values...)
	fmt.Println("Node 2 data:", head.Data)
	fmt.Println("\nShow Linked List:", head.PrintList())
	head.Append(94)
	fmt.Println("\nUse Append function:", head.PrintList())
	head.ChangeHead(78)
	fmt.Println("\nUse ChangeHead function:", head.PrintList())

	head.RemoveNode(235)
	fmt.Println("\nUse RemoveNode function:", head.PrintList())
	fmt.Println("Teste: ", head.CountNodes())
}

func FibonacciSequence() {
	fibSequence := fib.Fibonacci{Number: 10}.Recursive()
	fmt.Println(fibSequence)
}
