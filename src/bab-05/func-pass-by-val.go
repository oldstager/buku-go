package main

import "fmt"

func changeMe(theVal int) int {

	theVal++
	return theVal

}

func main() {

	theVal := 45

	fmt.Println(theVal)

	fmt.Println(changeMe(theVal))

	fmt.Println(theVal)

}
