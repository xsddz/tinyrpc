package main

import (
	"Advanced-Golang-Programming/ex/foobar"
	"fmt"
)

func init() {
	fmt.Println("init:main")
}

func main() {
	// theBugs.Run()
	// foobar.Sort([]int{1, 2})
	// kh.Run()

	// fmt.Println("123 + 0 = ", foobar.BigAdder("123", "0"))
	// fmt.Println("123 + 2 = ", foobar.BigAdder("123", "2"))
	// fmt.Println("123 + 22 = ", foobar.BigAdder("123", "22"))
	// fmt.Println("123 + 223 = ", foobar.BigAdder("123", "223"))
	// fmt.Println("123 + 2234 = ", foobar.BigAdder("123", "2234"))

	fmt.Println(foobar.BigMult("12345", "23"))
	fmt.Println(foobar.BigMult("23", "12345"))

	fmt.Println(foobar.BigMult("12345", "0"))
	fmt.Println(foobar.BigMult("0", "12345"))

	fmt.Println(foobar.BigMult("12345", "1"))
	fmt.Println(foobar.BigMult("1", "12345"))
}
