package main

import (
	"fmt"
	"os"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	stringPassed := os.Args[1]
	reversedString := Reverse(stringPassed)
	if stringPassed == reversedString {
		fmt.Println("Matches! It is a palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
