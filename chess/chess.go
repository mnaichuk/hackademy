package main

import "fmt"

func CheckInput(posA string, posB string) bool {
	if len(posA) != 2 || len(posB) != 2 {
		return false
	}
	if posA[0] < 97 || posA[0] > 122 || posA[1] < 49 || posA[1] > 56 {
		return false
	}
	if posB[0] < 97 || posB[0] > 122 || posB[1] < 49 || posB[1] > 56 {
		return false
	}
	return true
}

func CanKnightAttack(posA string, posB string) bool {
	if CheckInput(posA, posB) == false {
		return false
	}
	if (posA[0]-2) == posB[0] && (posA[1]+1) == posB[1] {
		return true
	} else if (posA[0]-2) == posB[0] && (posA[1]-1) == posB[1] {
		return true
	} else if (posA[0]+2) == posB[0] && (posA[1]-1) == posB[1] {
		return true
	} else if (posA[0]+2) == posB[0] && (posA[1]+1) == posB[1] {
		return true
	} else if (posA[0]-1) == posB[0] && (posA[1]-2) == posB[1] {
		return true
	} else if (posA[0]-1) == posB[0] && (posA[1]+2) == posB[1] {
		return true
	} else if (posA[0]+1) == posB[0] && (posA[1]-2) == posB[1] {
		return true
	} else if (posA[0]+1) == posB[0] && (posA[1]+2) == posB[1] {
		return true
	} else {
		return false
	}
}

func main() {
	posA := "f1"
	posB := "g2"
	res := CanKnightAttack(posA, posB)
	fmt.Println(res)
}
