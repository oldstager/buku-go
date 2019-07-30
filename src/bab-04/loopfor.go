package main

import "fmt"

func main() {
	var i, j int

	for i = 1; i < 10; i++ {
		fmt.Println(i)
		for j = 1; j <= i; j++ {
			fmt.Println(j)
		}
	}
}
