package main

import "fmt"

func array_sample() {
	//arr1 := [...]int{1, 2, 3, 4, 5} // #3
	// arr1 := [5]int{2: 1, 4: 5} // #4
	// arr5 := [...]int{2: 1, 4: 5}     // #5

	// 陣列大小
	// fmt.Println(arr1)
	// fmt.Println(arr1[1:3])

	// s2 := make([]int, 5)
	// fmt.Println(s2)

	// length := 5
	// res := make([]int, len(s2))
	// fmt.Println(res)

	dp := make([]int, 5)

	fmt.Println(dp)

	// arr1_len := len(arr1)
	// fmt.Println("array 1 length: ", arr1_len)
	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] += 1
	// 	fmt.Printf("[%d]: %d, ", i, arr1[i])
	// }
}

func getConcatenation(nums []int) []int {

	nums = append(nums, nums...)

	return nums
}

func main() {

	input := []int{1, 2, 3, 4, 5}
	result := getConcatenation(input)

	fmt.Println(result)

	// s := "abcde"
	// fmt.Println(s[0])
	// fmt.Println(s[1])
	// fmt.Println(rune(s[1]))
	// fmt.Println(string(s[0]))

	// array_sample()

	// s := "abc"

	// fmt.Println(len(s))
	// fmt.Println(s[0])
	// fmt.Println(s[0:1])
	// fmt.Println(string(s[0]))
	// fmt.Println(s[1])
	// fmt.Println(string(s[1]))         // ASCII only
	// fmt.Println(string([]rune(s)[0])) // UTF-8
	// fmt.Println(string([]rune(s)[1])) // UTF-8

	// s1 := "hello"
	// b := []byte(s1)
	// // []byte to string
	// s2 := string(b)

}
