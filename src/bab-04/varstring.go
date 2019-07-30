package main

import (
	"fmt"
	"reflect"
	s "strings"
)

// Definisi string
var str1 string = "UGM"
var str2 = "Yogyakarta"
var str3 = "universitas\ngadjah mada"

var str3backtick = `universitas\ngadjah mada`

// error: illegal rune literal
//var str3singlequoted = 'universitas gadjah mada'

func main() {

	// Lihat https://golang.org/pkg/strings/
	fmt.Println(str1)
	fmt.Println(len(str1))
	fmt.Println(s.Contains(str1, "GM"))
	fmt.Println(s.Title(str3))
	fmt.Println(str1[0])
	fmt.Println(s.Join([]string{str1, str2}, " "))
	fmt.Println(s.Join([]string{str3, str2}, "\n"))
	fmt.Println(s.Join([]string{str3backtick, str2}, "\n"))
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.TypeOf(str2))
	fmt.Println()

}
