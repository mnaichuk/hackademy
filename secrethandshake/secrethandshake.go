package main 

import "fmt"

func Handshake(n int) {
	arr := make([]byte,0)
	for ; n > 0; {
		arr = append(arr,byte(n%2))
		n = int(n / 2)
	}
	fmt.Println(arr)
}


func main() {
	Handshake(12)
}
