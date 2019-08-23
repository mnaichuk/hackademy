package main

import (
	"fmt"
)

func Decode(s string) (n int, ok bool) {
	i := 0
	n = 0
	if len(s) == 0 {
		return 0,false
	}
	for ; i < len(s) - 1; i++ {
		switch s[i] {
			case 'M':
				n = n + 1000
			case 'D':
				n = n + 500
			case 'C':
				if s[i + 1] == 'M' || s[i + 1] == 'D' {
					n = n - 100
				} else {
					n = n + 100
				}
			case 'L':
				n = n + 50
			case 'X':
				if s[i + 1] == 'L' || s[i + 1] == 'C' {
					n = n - 10
				} else {
					n = n + 10
				}
			case 'V':
				n = n + 5
			case 'I':
				if s[i + 1] == 'V' || s[i + 1] == 'X' {
					n = n - 1
				} else {
					n = n + 1
				}
			default:
				return 0,false
		}
	}
	switch s[i] {
		case 'M':
			n = n + 1000
		case 'D':
			n = n + 500
		case 'C':
			n = n + 100
		case 'L':
			n = n + 50
		case 'X':
			n = n + 10
		case 'V':
			n = n + 5
		case 'I':
			n = n + 1
		default:
			return 0,false
	}
	return n,true
}



func main() {
	fmt.Println(Decode("I"))
	fmt.Println(Decode(""))
	fmt.Println(Decode("1"))
	fmt.Println(Decode("T"))
	fmt.Println(Decode("X X X"))
	fmt.Println(Decode("X0"))
	fmt.Println(Decode("XVIII"))
	fmt.Println(Decode("XXXVIII"))
	fmt.Println(Decode("CX"))
	fmt.Println(Decode("CXXXVIII"))
	fmt.Println(Decode("CCCLXVI"))
	fmt.Println(Decode("DCXXXIV"))
	fmt.Println(Decode("MMDCCXIV"))
	fmt.Println(Decode("MMMCD"))
		
}
