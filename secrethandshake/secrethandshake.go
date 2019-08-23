package main

import "fmt"

func Handshake(i int) []string {
	var answer []string
	if i&1 == 1 {
		answer = append(answer, "wink")
	}
}
