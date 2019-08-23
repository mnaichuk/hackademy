package main

import "fmt"

func CanAttack(p1, p2 string) bool {
	if p1[0]+2 == p2[0] {
		if p1[1]+1 == p2[1] {
			return true
		}
		if p1[1]-1 == p2[1] {
			return true
		}
	}
	if p1[0]-2 == p2[0] {
		if p1[1]+1 == p2[1] {
			return true
		}
		if p1[1]-1 == p2[1] {
			return true
		}
	}
	if p1[1]+2 == p2[1] {
		if p1[0]+1 == p2[0] {
			return true
		}
		if p1[0]-1 == p2[0] {
			return true
		}
	}
	if p1[1]-2 == p2[1] {
		if p1[0]+1 == p2[0] {
			return true
		}
		if p1[0]-1 == p2[0] {
			return true
		}
	}
	return false
}

func main() {
	can := CanAttack("g6", "e4")
	fmt.Println(can)
}
