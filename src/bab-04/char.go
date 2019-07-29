/*

	diambil dari: https://www.callicoder.com/golang-basic-types-operators-type-conversion/

*/
package main

import "fmt"

func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}
