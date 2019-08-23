package main

import (
	"fmt"
)

func repStr(s string, n int) (res string) {
	res = ""
	for i := 0; i < n; i++ {
		res = res + s
	}
	return res
}

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

func Encode(n int) (s string, ok bool) {
	if n < 0 {
		return "",false
	}
	s = ""
	for ; n >= 1000; {
		n = n - 1000
		s = s + "M"
	}
	if n >= 900 {
		s = s + "CM"
		n = n - 900
	}
	if n >= 500 {
		s = s + "D"
		n = n - 500
	}
	if n >= 400 {
		s = s + "CD"
		n = n - 400
	}
	k := int(n / 100)
	n = n - 100*k
	s = s + repStr("C",k)
	if n >= 90 {
		n = n - 90
		s = s + "XC"
	}
	if n >= 50 {
		n = n - 50
		s = s + "L"
	}
	if n >= 40 {
		n = n - 40
		s = s + "XL"
	}
	k = int(n / 10)
	s = s + repStr("X",k)
	n = n - k*10
	if n == 9 {
		return s + "IX",true
	}
	if n >= 5 {
		n = n - 5
		s = s + "V"
	}
	if n == 4 {
		return s + "IV",true
	}
	return s + repStr("I",n),true
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

	fmt.Println("\n  Encode test \n")
	fmt.Println(Encode(3943))
}
