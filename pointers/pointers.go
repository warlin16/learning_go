package pointers

import "fmt"

// PrintAddress prints the address of an int
func PrintAddress() {
	x := 16
	fmt.Println("the value of x", x)
	// you use the & symbol to see the address of a value
	fmt.Println("the address of x", &x)
	fmt.Println("**********************")
	fmt.Printf("%T\n", x)  // <- print the type which is int
	fmt.Printf("%T\n", &x) // <- *int a *POINTER* to int
	fmt.Println("**********************")
	// using pointers as a value for a variable
	y := &x // <- assigning the *POINTER* to y
	fmt.Println("the pointer:", y)
	// but to see the value stored at that address you use the * symbol
	fmt.Println("and it's value:", *y)
	// since you have a pointer to the address for the value of x
	// you can change that value with y
	*y = 27
	fmt.Println("Changed the value of x by using it's pointer.. wow", x)
}
