package main

import "fmt"

func findSumOfChars(strCounted string, theChar rune) int {
	counter := 0
	for _, c := range strCounted {
		if c == theChar {
			counter++
		}
	}
	return counter
}

func main() {
	theStr := "STMIK Akakom"
	fmt.Printf("Jumlah karakter 'k' dari string %s adalah %d",
		theStr, findSumOfChars(theStr, 'k'))
}
