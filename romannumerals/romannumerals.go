package main

func Encode(number int) (string, bool) {
	convert := []struct {
		value int
		digit string
	}{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}

	roman := ""
	for _, c := range convert {
		for number >= convert.value {
			roman += convert.digit
			number -= convert.value
		}
	}
	return roman, true
}

var rom = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		vchar := rom[char]
		if i < len(str)-1 {
			cnext := string(str[i+1])
			vcnext := rom[cnext]
			if vchar < vcnext {
				res += vcnext - vchar
				i++
			} else {
				res += vchar
			}
		} else {
			res += vchar
		}
	}
	return res
}

func main() {
	//when input is "", the function won't be called!!!
}
