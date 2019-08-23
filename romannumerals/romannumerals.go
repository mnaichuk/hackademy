package main

import "fmt"

var Rome = map[int64]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func Encode(num int64) (string, bool) {
	var str string
	var a int64 = 0

	if num <= 0 {
		return "", false
	}
	for num > 0 {
		if num >= 1000 {
			a = 1000
		} else if num >= 900 {
			a = 900
		} else if num >= 500 {
			a = 500
		} else if num >= 400 {
			a = 400
		} else if num >= 100 {
			a = 100
		} else if num >= 90 {
			a = 90
		} else if num >= 50 {
			a = 50
		} else if num >= 40 {
			a = 40
		} else if num >= 10 {
			a = 10
		} else if num >= 9 {
			a = 9
		} else if num >= 5 {
			a = 5
		} else if num >= 4 {
			a = 4
		} else if num >= 1 {
			a = 1
		}
		str += Rome[a]
		num -= a
	}
	return str, true
}

func Decode(num string) (int64, bool) {
	var count int64 = 0

	if len(num) < 1 {
		return 0, false
	}
	for i := range num {
		switch num[i] {
		case 'M':
			count += 1000
		case 'D':
			if i > 0 {
				if num[i-1] == 'C' {
					count += 300
				} else {
					count += 500
				}
			} else {
				count += 500
			}
		case 'C':
			if i > 0 {
				if num[i-1] == 'X' {
					count += 80
				} else {
					count += 100
				}
			} else {
				count += 100
			}
		case 'L':
			if i > 0 {
				if num[i-1] == 'X' {
					count += 30
				} else {
					count += 50
				}
			} else {
				count += 50
			}
		case 'X':
			if i > 0 {
				if num[i-1] == 'I' {
					count += 8
				} else {
					count += 10
				}
			} else {
				count += 10
			}
		case 'V':
			if i > 0 {
				if num[i-1] == 'I' {
					count += 3
				} else {
					count += 5
				}
			} else {
				count += 5
			}
		case 'I':
			count += 1
		default:
			return 0, false
		}
	}
	return count, true
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

	fmt.Println(Encode(-1))

	fmt.Println(Encode(0))
	fmt.Println(Encode(1))
	fmt.Println(Encode(18))
	fmt.Println(Encode(38))
	fmt.Println(Encode(110))
	fmt.Println(Encode(138))
	fmt.Println(Encode(366))
}
