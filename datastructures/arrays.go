package datastructures

import "fmt"

var myArr [5]int

// PrintArr prints stuff
func PrintArr() {
	myArr[0] = 1
	myArr[1] = 2
	myArr[2] = 3
	myArr[3] = 4
	myArr[4] = 5
	fmt.Println("print this array bro", myArr)
}
