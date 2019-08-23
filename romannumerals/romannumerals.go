package main

import "fmt"

func Encode(num int) (string, bool) {
	if num <= 0 {
		return "", false
	}
	var roman string
	for num > 0 {
		switch {
		case num >= 1000:
			roman += "M"
			num -= 1000
		case num >= 900:
			roman += "CM"
			num -= 900
		case num >= 500:
			roman += "D"
			num -= 500
		case num >= 400:
			roman += "CD"
			num -= 400
		case num >= 100:
			roman += "C"
			num -= 100
		case num >= 90:
			roman += "XC"
			num -= 90
		case num >= 50:
			roman += "L"
			num -= 50
		case num >= 40:
			roman += "XL"
			num -= 40
		case num >= 10:
			roman += "X"
			num -= 10
		case num >= 9:
			roman += "IX"
			num -= 9
		case num >= 5:
			roman += "V"
			num -= 5
		case num >= 4:
			roman += "IV"
			num -= 4
		case num >= 1:
			roman += "I"
			num -= 1
		}
	}
	return roman, true
}

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(rom string) int {
	dec := 0
	for i := 0; i < len(rom); i++ {
		char := num[string(rom[i])]
		if i < len(rom)-1 {
			nextchar := num[string(rom[i+1])]
			if char < nextchar {
				dec += nextchar - char
				i++
			} else {
				dec += char
			}
		} else {
			dec += char
		}
	}
	return dec
}

func main() {
	str, ok := Encode(2714)
	fmt.Println(Decode(str))
	fmt.Println(ok)
	fmt.Println(str)
}
