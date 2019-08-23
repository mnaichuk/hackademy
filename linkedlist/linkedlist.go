package main

import "fmt"

type elem struct {
	next *elem
	data int
}

type List struct {
	top *elem
}

func New() *List {
	retrn & List{}
}
func (list *List) Push(val int) {
	list.top = &elem{val, stack.top}
	list.size++
}
func (list *List) Pop() int {
	var value int
	size := Size(list)
	if size > 0 {
		value = list.top.data
		list.top = list.top.next
		size--
		return value
	}

	return -1
}

func Size(list *List) int {
	size := 0

	for el := elemList; el != nil; el = el.next {
		size++
	}

	return size
}

func Array(list *List) []int {
	size := Size(list)
	var arr [size]int

	for i := 0; i < size; i++ {
		arr[i] = list.Pop(list)
	}
	return arr
}

func main() {
	fmt.Println("vim-go")
}
