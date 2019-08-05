// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Jika cat tidak mempunyai Speak, maka akan muncul error:
// ./cat.go:27:32: cannot use Cat literal (type Cat) as type Animal in array or slice literal:
//	Cat does not implement Animal (missing Speak method)

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
