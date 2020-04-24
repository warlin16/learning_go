package functions

import "fmt"

// PrintExpressions prints func expressions
func PrintExpressions() {
	// you can assign funcs to variables, similar to JS
	f := func() {
		fmt.Println("I've been assigned to a variable even though I started off as an anon func!!")
	}
	f()
	fmt.Println("---------------------------------------")
	surname := myNameAndReturnLast()
	fmt.Println("now this is my last name:", surname())
	fmt.Println("---------------------------------------")
	printSumOfNums()
	fmt.Println("---------------------------------------")
	a := incrementor()
	a()
	a()
	a()
	a()
	a()
	fmt.Println("The value of a", a())
	b := incrementor()
	fmt.Println("The value of b", b())
}

// returning a func from a func
func myNameAndReturnLast() func() string {
	fmt.Println("My name is warlin and you'll have to call the return value...")
	return func() string {
		return "Reyes"
	}
}

// ************************* Callbacks, closures **************************************
// make a func sum and supply that as a call back to a func sumEven and sumOdd
// sumEven sums all even numbers while sumOdd sums all odd numbers

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sumEven(cb func(nums ...int) int, nums ...int) int {
	var evenNums []int
	for _, num := range nums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}
	return cb(evenNums...)
}

func sumOdd(cb func(nums ...int) int, nums ...int) int {
	var oddNums []int
	for _, num := range nums {
		if !(num%2 == 0) {
			oddNums = append(oddNums, num)
		}
	}
	return cb(oddNums...)
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // <- package level scope

func printSumOfNums() {
	fmt.Println("This is the sum of all even numbers in nums", sumEven(sum, nums...))
	fmt.Println("This is the sum of all odd numbers in nums", sumOdd(sum, nums...))
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
