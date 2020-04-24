package functions

import "fmt"

// PrintFunctions prints functions
func PrintFunctions() {
	defer funcToDefer()
	defer secondDefer()
	printName("warlin")
	fmt.Println("Warlin's last name is", returnLastName("Reyes"))
	variadicParams(1, 3, 5, 7, 9)
	// since we know the above function accepts variadic params
	// we can destructure a slice and pass it in as an argument
	xi := []int{1, 2, 3, 4, 5}
	variadicParams(xi...)
}

func printName(s string) {
	fmt.Println("Hello,", s)
}

func returnLastName(str string) string {
	return str
}

// you can define a slice of type as a parameter for functions
// called a variadic parameter
func variadicParams(x ...int) {
	fmt.Println("This is probably going to be a slice", x)
}

// Defer is is pushing a functions execution to the end of the function it
// is called in, for example if you open a file you can defer the function
// that closes that file once the code is done executing
func funcToDefer() {
	fmt.Println("I'm going to print at the end even though I'm being called first.")
}

func secondDefer() {
	fmt.Println("Wondering if defer works like a queue or stack. The answer is a stack")
}
