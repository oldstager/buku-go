package main

import "fmt"

var theVal int = 0

func changeMe(theVal *int) int {

	*theVal = 150

	return *theVal

}

func main() {

	theVal := 45

	fmt.Println(theVal)

	fmt.Println(changeMe(&theVal))

	fmt.Println(theVal)

}
