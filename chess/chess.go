package main

import "fmt"

func CanKnightAttack(x, y string) bool {

	if x[0]+2 == y[0] && x[1]+1 == y[1] {
		return true

	} else if x[0]+2 == y[0] && x[1]-1 == y[1] {
		return true

	} else if x[0]-2 == y[0] && x[1]+1 == y[1] {
		return true

	} else if x[0]-2 == y[0] && x[1]-1 == y[1] {
		return true

	} else if x[1]+2 == y[1] && x[0]+1 == y[0] {
		return true

	} else if x[1]+2 == y[1] && x[0]-1 == y[0] {
		return true

	} else if x[1]-2 == y[1] && x[0]+1 == y[0] {
		return true

	} else if x[1]-2 == y[1] && x[0]-1 == y[0] {
		return true
	}

	return false
}

func main() {
	pos := CanKnightAttack("g5", "e4")
	fmt.Println(pos)
}
