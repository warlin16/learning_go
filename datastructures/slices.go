package datastructures

import "fmt"

// a SLICE allows you to group values of the same TYPE

// PrintSlice prints shit br0
func PrintSlice() {
	x := []int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println("this is the index and the value", i, v)
	}
	multiDimensionalSlice()
}

func slicingASlice() {
	x := []int{12, 13, 14, 15}
	// slices from position 1 until the end
	fmt.Println(x[1:])
	// slice from positon 1 up to but not including position 3
	fmt.Println(x[1:3])
	// slice from the beginning of the slice up to but not including position 2
	fmt.Println(x[:2])
}

func appendingToSlice() {
	// appending to a slice you have to use the append() function
	// the append function returns a new slice IT DOES NOT MODIFY THE ORIGINAL
	x := []string{"warlin", "reyes"}
	x = append(x, "and", "ciera", "romero")
	fmt.Println(x)
	// using the ... operator after a type works similar to how it does in JS except you
	// have to use it after the variable
	y := []string{"and", "they", "want", "a", "dog"}
	x = append(x, y...)
	fmt.Println("coding is so cool bro", x)
}

func deletingFromSlice() {
	x := []int{1, 2, 7, 6, 3, 4, 5}
	x = append(x[:2], x[4:]...)
	fmt.Println("this is the correct slice", x)
}

func makeASlice() {
	// be wary of the cap of a slice, appending to a slice and going over the cap
	// causes it to make a new underlying array double the size and copies all
	// the values of the old array and throws the old array away using processing power
	x := make([]int, 2, 3)
	fmt.Println(x)
	x = append(x, 1)
	fmt.Println(x)
	fmt.Println(cap(x))
	fmt.Println("the slice should double in size now")
	x = append(x, 2)
	fmt.Println(x, cap(x))

	// curious to see what the cap is when you don't specify it
	y := []int{1, 2, 3}
	fmt.Println("this is the new slice y and its cap", y, cap(y))
	// so using composite literals to make a slice makes it's cap whatever amount of elements
	// you supply to it when you make it the first time
}

func multiDimensionalSlice() {
	names := []string{"warlin", "ciera"}
	surnames := []string{"reyes", "romero"}
	multiDimensional := [][]string{names, surnames}
	fmt.Println("the 2d slice", multiDimensional)
}
