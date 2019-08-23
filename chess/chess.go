package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func CanKnightAttack(a,b string) bool {
	x1 := int(a[0] - 'a' + 1)
	y1 := int(a[1] - '0')
	x2 := int(b[0] - 'a' + 1)
	y2 := int(b[1] - '0')
	if x1 > 8 || x1 < 1 || x2 < 1 || x2 > 8 || y1 > 8 || y1 < 1 || y2 > 8 || y2 < 1 {
		return false
	}
	h1 := abs(x1 - x2)
	h2 := abs(y1 - y2)
	if (h1 == 1 && h2 == 2) || (h1 == 2 && h1 == 1) {
		return true
	}
	return false
}

func main() {
	fmt.Println(CanKnightAttack("a2","b4"))
}
