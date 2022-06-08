package main

import "fmt"

func main() {
	// s := "abc"

	// fmt.Println(len(s))
	// fmt.Println(s[0])
	// fmt.Println(s[0:1])
	// fmt.Println(string(s[0]))
	// fmt.Println(s[1])
	// fmt.Println(string(s[1]))         // ASCII only
	// fmt.Println(string([]rune(s)[0])) // UTF-8
	// fmt.Println(string([]rune(s)[1])) // UTF-8

	s1 := "hello"
	b := []byte(s1)
	// []byte to string
	s2 := string(b)

	fmt.Println(s1)
	fmt.Println(b)
	fmt.Println(s2)
}
