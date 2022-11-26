package main

import "fmt"

func main() {

	test := []byte{'.', '.', '.'}

	fmt.Println(test)

	fmt.Println(string(test[0]))

	test[0] = '#'

	fmt.Println(test)

	fmt.Println(string(test[0]))

	// obj := []int{1, 2, 3, 4, 5}

	// obj2 := 10

	// fmt.Println(obj)

	// test(obj)
	// test2(obj2)

	// fmt.Println(obj)
	// fmt.Println(obj2)

	// obj3 := []int{1, 2, 3, 4, 5}

	// test3(&obj3)

	// fmt.Println(obj3)
}

func test(items []int) {
	modify(items)
}

func modify(tmp []int) {
	tmp[2] = 99
}

func test2(item int) {
	modify2(item)
}

func modify2(tmp int) {
	tmp = 3
}

func test3(items *[]int) {
	(*items)[3] = 100
	fmt.Println(len((*items)))
}
