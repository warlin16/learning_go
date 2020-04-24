package functions

import "fmt"

// going to test adding methods to structs
type person struct {
	first, last string
}

type employee struct {
	person
	jobTitle string
}

func (p person) myName() {
	fmt.Println("My name is", p.first, p.last)
}

func (e employee) myName() {
	fmt.Println("Hey I am", e.first, e.last, "and I'm a", e.jobTitle)
}

func printHunan(h human) {
	switch h.(type) {
	case person:
		fmt.Println("This human is of type person")
	case employee:
		fmt.Println("This human is of type employee")
	}
	h.myName()
}

// PrintMethod prints a method
func PrintMethod() {
	warlin := person{
		"warlin",
		"reyes",
	}
	fmt.Println("this is", warlin)
	printHunan(warlin)
	fmt.Println("a person is a human but an employee can also be a human, check it out")
	fmt.Println("******************************")
	reyes := employee{
		person{
			"Warlin Felipe",
			"Reyes",
		},
		"Software Engineer",
	}
	printHunan(reyes)
	fmt.Println("----------------------- Testing empty interfaces as params -----------------------")
	printAnything(16)
	printAnything("Same function above me ran with an int... wowza!")
	// anonymous functions are similar to JS
	func() {
		fmt.Println("Im an anon func")
	}()
}

// interfaces :eyes:
// using an interface adds a second type to which ever struct
// that uses the function the interface defines
type human interface {
	myName()
}

// Every single type implements an empty interface
func printAnything(element ...interface{}) {
	fmt.Println("We can print anything", element[0])
}
