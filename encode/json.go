package encode

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
}

// Test tests
func Test() {
	p := person{
		"warlin",
		"reyes",
	}
	p2 := person{
		"ciera",
		"romero",
	}

	ppl := []person{p, p2}

	pplJSON, err := json.Marshal(ppl)
	if err != nil {
		fmt.Println("This is the error:", err)
	}
	fmt.Println("This is the json", string(pplJSON))
}

type things struct {
	Hobby string `json:"Hobby"`
}

// OtherTest unmarshaling json
func OtherTest() {
	js := `{"Hobby": "Go Programmer"}`
	// the unmarshal function takes a byte slice argument so we have to convert
	bs := []byte(js)
	var hobby things
	err := json.Unmarshal(bs, &hobby)
	if err != nil {
		fmt.Println("This is the error", err)
	}
	fmt.Println("the data", hobby)
}
