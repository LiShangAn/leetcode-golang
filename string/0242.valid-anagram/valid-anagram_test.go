package leetcode

import (
	"fmt"
	"testing"
)

type question242 struct {
	para242
	ans242
}

type para242 struct {
	one string
	two string
}

type ans242 struct {
	result bool
}

func Test_Question242(t *testing.T) {

	qs := []question242{

		{
			para242{"anagram", "nagaram"},
			ans242{true},
		},

		{
			para242{"rat", "car"},
			ans242{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans242, q.para242
		fmt.Printf("input:%v output:%v\n", p, isAnagram(p.one, p.two))
	}
}
