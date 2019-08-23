package main

import "fmt"

var numRom = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
var numInt = map[int]string{
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

var table = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func biggestNumber(num int) int {
	for _, tmp := range table {
		if tmp <= num {
			return tmp
		}
	}

	return 1
}

func Encode(decint int) (string, bool) {
	res := ""
	if decint < 1 {
		return res, false
	}
	for decint > 0 {
		tmp := biggestNumber(decint)
		res += numInt[tmp]
		decint -= tmp
	}

	return res, true
}

func Decode(romnum string) (int, bool) {
	res := 0
	if romnum == "" {
		return 0, false
	}
	lenght := len(romnum)
	for i := 0; i < lenght; i++ {
		tmp := string(romnum[i])
		dnum := numRom[tmp]
		if i < lenght-1 {
			tmpnext := string(romnum[i+1])
			dnumnext := numRom[tmpnext]
			if dnum < dnumnext {
				res += dnumnext - dnum
				i++
			} else {
				res += dnum
			}
		} else {
			res += dnum
		}
	}

	return res, true
}

func main() {

	str := "MCC"
	dec, decvalid := Decode(str)
	fmt.Printf("%s = %d %v\n", str, dec, decvalid)

	tmp := 17
	enc, encvalid := Encode(tmp)
	fmt.Printf("%d = %s %v", tmp, enc, encvalid)

}
