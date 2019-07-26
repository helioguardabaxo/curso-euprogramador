package main

import "fmt"

func main()  {
	b := true
	number := 2

	if b {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if n := 5; n > 0 {
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor que 0")
	}

	switch number {
	case 2:
		fmt.Println("Number 2")
	case 5:
		fmt.Println("Number 5")
	default:
		fmt.Println("Default")
	}
}