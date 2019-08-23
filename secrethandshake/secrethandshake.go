package main

import "fmt"

var moves = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(i int) []string {

	var res []string

	if i&1 == 1 {
		res = append(res, moves[0])
	}

	if i&2 == 2 {
		res = append(res, moves[1])
	}

	if i&4 == 4 {
		res = append(res, moves[2])
	}

	if i&8 == 8 {
		res = append(res, moves[3])
	}
	if i&16 == 16 {
		for j := 0; j < len(res)/2; j++ {
			tmp := res[j]
			res[j] = res[len(res)-1-j]
			res[len(res)-1-j] = tmp
		}
	}

	return res
}

func main() {
	input := 19
	n := Handshake(input)
	fmt.Println(n)
	fmt.Println("\nvim-go")
}
