package controlflow

import "fmt"

// LoopAndPrintASCII bro
func LoopAndPrintASCII() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
