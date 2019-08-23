package main

import "fmt"

type LinkedList struct {
	size int
	Top  *Node
}

type Node struct {
	Data int
	Next *Node
}

func New(data int) *LinkedList {
	return &LinkedList{1, &Node{data, nil}}
}

func (list *LinkedList) Push(data int) {
	tmp := &Node{data, list.Top}
	list.size++
	list.Top = tmp
}

func (list *LinkedList) Pop() int {
	if list.size > 0 {
		data := list.Top.Data
		list.Top = list.Top.Next
		list.size--
		return data
	} else {
		return 0
	}
}

func (list *LinkedList) Array() []int {
	arr := make([]int, list.size)
	tmp := list.Top
	for i := 0; i < list.size; i++ {
		arr[i] = tmp.Data
		if tmp.Next != nil {
			tmp = tmp.Next
		}
	}
	return arr
}

func (list *LinkedList) Reverse() {
	data := list.Pop()
	fmt.Println(data)
	newlist := New(data)
	fmt.Println(newlist)
	for list.size > 0 {
		data = list.Pop()
		newlist.Push(data)
	}
	list.size = newlist.size
	list.Top = newlist.Top
}

func (list *LinkedList) Size() {
	fmt.Println("Size: ", list.size)
}

func main() {
	fmt.Println("vim-go")
	list := New(3)
	list.Push(9)
	list.Push(10)
	list.Reverse()
	fmt.Println(list)
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
}
