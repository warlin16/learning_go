package datastructures

import "fmt"

type person struct {
	first, last string
}

// this is called embedding when you use another struct as part of
// another structs definition
type hispanic struct {
	person
	citizen bool
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

	reyes := hispanic{
		person: person{
			first: "Warlin",
			last:  "Reyes",
		},
		citizen: true,
	}

	if reyes.citizen {
		fmt.Println(reyes.last, "You are a hispanic citizen!")
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
