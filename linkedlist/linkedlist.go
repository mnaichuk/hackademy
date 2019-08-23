package main

import "fmt"

type (

	Stack struct {
	top *Item
	size int
	}


	Item struct {
		value int
		next *Item 
	}

)

func NewStack() *Stack {
 return &Stack{}
}

func (s *Stack) Push (val int) {
	s.top = &item(value, s.top)
	s.size++
}

func (s *Stack) Pop () (val int)  {
	if s.size > 0 {
		val := s.top.val
		s.top = s.top.next
		s.size--
		return val
	}
	return 1
}

func (s *Stack) ToArr() ([int]int) {
	val arr [s.size]int 
	tmp := s.top
	for i := 0; i < s.size; i++ {
		arr[i] = tmp.val
		if s.top.next != nil {
		tmp = s.top.next
		}	
	}
	return arr[s.size]
}


func FromArr (a [int]int) (*Stack){
	st := NewStack()
	for i := 0; i < n; i++ {
		st.top.value = a[i]
		if i == n - 1 {
		st.top.next = nil
		}		
	}
	return st
}

func (s *Stack) Reverse() (*Stack) {
	arr = s.ToArr
	for int i = 0; i < len(arr); i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr) - i]
		arr[len(arr) - i] = tmp
	}
	newStack = FromArr(arr)
	return newStack
}

