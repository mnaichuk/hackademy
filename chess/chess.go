package main

import "fmt"

func CanKnightAttack(pos1, pos2 string) bool {
	p11 := pos1[0]
	p12 := pos1[1]
	p21 := pos2[0]
	p22 := pos2[1]
	if p11 >= 97 && p11 <= 104 && p21 >= 97 && p21 <= 104 && p12 >= 49 && p12 <= 56 && p22 >= 49 && p22 <= 56 {
		if p11+2 == p21 && p12+1 == p22 {
			return true
		} else if p11+1 == p21 && p12+2 == p22 {
			return true
		} else if p11-1 == p21 && p12+2 == p22 {
			return true
		} else if p11-2 == p21 && p12+1 == p22 {
			return true
		} else if p11-2 == p21 && p12-1 == p22 {
			return true
		} else if p11-1 == p21 && p12-2 == p22 {
			return true
		} else if p11+1 == p21 && p12-2 == p22 {
			return true
		} else if p11+2 == p21 && p12-1 == p22 {
			return true
		}
	} else {
		fmt.Printf("invalid position\n")
		return false
	}

	return false
}

func main() {

	str1 := "a8"
	str2 := "i6"
	res := CanKnightAttack(str1, str2)
	fmt.Printf("%s %s %v", str1, str2, res)
}
