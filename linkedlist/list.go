package main

import "fmt"

type List struct {
	data string
	next *List
}

type Stack struct {
	head    *List
	counter int
}

func New() *Stack {
	return &Stack{counter: 0}
}

func (s *Stack) Push(Ndata string) {
	var list *List

	list.data = Ndata
	list.next = s.head
	s.head = list
	s.counter++
}

func (s *Stack) Pop() string { // Pop returns list data
	pop := s.head.data
	s.head = s.head.next
	s.counter--
	return pop
}

func (s *Stack) Array(arr []string) { //Method takes array of data and push this data (inside list) to the stack
	l := len(arr)
	if l < 1 {
		return
	}
	for i := range arr {
		s.Push(arr[i])
	}
}

func (s *Stack) Reverse() *Stack {
	NewStack := New()
	for j := s.counter; j >= 0; j-- {
		NewStack.Push(s.Pop())
	}
	return NewStack
}

func (s *Stack) Size() int {
	return s.counter
}

func main() {
	fmt.Println("vim-go")
}
