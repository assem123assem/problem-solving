package main

import (
	"fmt"
	"problem-solving/problem1"
	"problem-solving/problem3"
	"problem-solving/problem4"
	"problem-solving/problem5"
	"problem-solving/problem6"
	"problem-solving/problem8"
	"problem-solving/problem9"
)

func sol1() {
	var word = "aasaa"
	fmt.Println(problem1.IsPalindrome(word))
}

func sol3() {
	fmt.Println(problem3.RunLengthEncode("xxxzzzxxaaaaacz"))
	fmt.Println(problem3.RunLengthDecode("x3z3x2a5c1z1"))
}

func sol4() {
	composedFunction := problem4.Compose(problem4.Square, problem4.AddOne)
	fmt.Println(composedFunction(6))
}

func sol5() {
	names := []string{"ahmed", "assem", "assem", "ali"}
	fmt.Println(problem5.Unique(names))
}

func sol6() {
	var original problem6.Matrix = [][]int{
		{0, 1, 2},
		{4, 5, 6},
		{8, 9, 10},
	}
	fmt.Println("original", original)
	transposedMat := problem6.Transpose(original)
	fmt.Println("transposed", transposedMat)
}

func sol8() {
	var arr = []int{1, 2, 3, 5, 6, 2, 4, 5}
	index := problem8.FirstDuplicateIndex(arr)
	if index == -1 {
		fmt.Println(arr, " has not duplicates")
	} else {
		fmt.Println("index of first duplicate is", index)
	}

}

func sol9() {
	tree1 := problem9.NewNode(1, []problem9.Node{
		problem9.NewNode(1, []problem9.Node{}),
		problem9.NewNode(2, []problem9.Node{}),
	})

	tree2 := problem9.NewNode(1, []problem9.Node{
		problem9.NewNode(1, []problem9.Node{}),
		problem9.NewNode(1, []problem9.Node{}),
	})

	// expected result =>  4
	fmt.Println(problem9.TreeSum(&tree1, 0))
	// expected result =>  3
	fmt.Println(problem9.TreeSum(&tree2, 0))

}

func main() {
	sol1()
	sol3()
	sol4()
	sol5()
	sol6()
	sol8()
	sol9()

}
