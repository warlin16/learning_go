package datastructures

import "fmt"

type person struct {
	first, last string
}

// this is called embedding when you use another struct as part of
// another structs definition
type employee struct {
	person
	employed bool
}

// PrintStruct prints structs
func PrintStruct() {
	warlin := person{
		first: "Warlin",
		last:  "Reyes",
	}
	// just like objects in JS you can access the fields on the struct
	// with . notation
	fmt.Println("My first name is", warlin.first, "and my last name is", warlin.last)

	reyes := employee{
		person: person{
			first: "Warlin",
			last:  "Reyes",
		},
		employed: true,
	}

	if reyes.employed {
		fmt.Println(reyes.last, "You are a employed!")
	}

	anonymousStruct()
}

func anonymousStruct() {
	car := struct {
		make, model string
		year        int
	}{
		make:  "Suzuki",
		model: "Kizashi",
		year:  2012,
	}
	fmt.Println("My name is warlin and I drive a", car.year, car.make, car.model)
}
