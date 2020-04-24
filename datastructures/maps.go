package datastructures

import "fmt"

// PrintMap prints maps
func PrintMap() {
	m := map[string]int{
		"Warlin": 27,
		"Ciera":  27,
	}
	fmt.Println("This is your map", m)
	// comma ok idiom
	if value, ok := m["Sarah"]; ok {
		fmt.Println("We should not be able to print the value", value)
	} else {
		fmt.Println("This is a common way to check maps for values")
	}
	m["Sarah"] = 27
	if value, ok := m["Sarah"]; ok {
		fmt.Println("We now added Sarah to our map, yay!", value, m)
	}
	fmt.Println("****************************")
	deleteFromMap()
}

func rangeOverMap() {
	m := map[string]int{
		"warlin": 16,
		"reyes":  15,
	}
	for k, v := range m {
		fmt.Println("These are the key value pairs:", k, v)
	}
}

func deleteFromMap() {
	m := map[string]int{
		"warlin": 1,
		"felipe": 2,
		"reyes":  3,
	}
	fmt.Println("the map with warlin in it", m)
	delete(m, "warlin")
	fmt.Println("the map without warlin in it", m)
	// interestingly enough you could delete a non-existent entry and get no error
}
