package golang

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	data [][]byte
}

type ans1 struct {
	result1 []int
}

func Test_Question1(t *testing.T) {
	qs := []question1{
		{
			para1{
				[][]byte{
					{'.', '.', '.'},
					{'.', '.', '.'},
					{'.', '.', '.'},
				}},
			ans1{[]int{0, 0, 0}},
		},
		{
			para1{
				[][]byte{
					{'.', '#', '.', '.'},
					{'.', '.', '#', '#'},
					{'.', '.', '.', '.'},
					{'#', '#', '#', '.'},
				}},
			ans1{[]int{1, 1, 1}},
		},
		{
			para1{
				[][]byte{
					{'#', '#', '.', '.', '.', '#'},
					{'#', '.', '.', '#', '.', '#'},
					{'.', '#', '#', '.', '.', '#'},
				}},
			ans1{[]int{1, 1, 2}},
		},
	}

	for _, q := range qs {
		input := q.para1
		expect := q.ans1.result1

		actual := countBoat(input.data)
		fmt.Printf("【expect】:%v\n【actual】:%v\n", expect, actual)
	}
}
