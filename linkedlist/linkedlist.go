package main

import (
	"fmt"
)

type Elem struct {
	data int
	next *Elem
}

type List struct {
	head *Elem
	top *Elem
	size int
}

func New() *List {
	list := List{nil,nil,0}
	return &list
}

func (l *List) push(data int) {
	elem := &Elem{data,nil}
	if l.size == 0 {
		l.top = elem
		l.head = elem
	} else {
		l.top.next = elem
		l.top = elem
	}
	l.size += 1
}

func (l *List) pop() (int,bool) {
	if l.size == 0 {
		return 0,false
	}
	data := l.top.data
	if l.size == 1 {
		l.top = nil
		l.head = nil
		l.size = 0
		return data,true
	}
	elem := l.head
	for i := 0; i < l.size - 2; i++ {
		elem = elem.next
	}
	elem.next = nil
	l.top = elem
	l.size -= 1
	return data,true
}

func (l *List) Size() int {
	return l.size
}

func (l *List) toArray() []int {
	arr := make([]int,0)
	elem := l.head
	for i:=0; i < l.Size(); i++ {
		arr = append(arr,elem.data)
		elem = elem.next
	}
	return arr
}

func ListFromArray(arr []int) *List {
	l := New()
	for i := 0; i < len(arr); i++ {
		l.push(arr[i])
	}
	return l
}

func (l *List) Reverse() *List {
	arr := l.toArray()
	lenn := len(arr)
	for i:=0; i < lenn/2; i++ {
		arr[i],arr[lenn - 1 - i] = arr[lenn - 1 - i], arr[i]
	}
	return ListFromArray(arr)
}

func main() {
	l := New()
	l.push(1)
	l.push(2)
	l.push(3)
	l.Reverse()
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println("Size: ",l.Size())
	fmt.Println(l.toArray())
}
