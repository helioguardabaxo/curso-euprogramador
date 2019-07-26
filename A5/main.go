package main

import "fmt"

func main()  {
	tabuada := [10]int{0, 5, 10}

	tabuada[3] = 15 

	fmt.Println(tabuada)
	fmt.Println(len(tabuada))

	usuario := map[string] string{
		"name" : "Helio",
		"nick" : "Eu mesmo",
	}

	usuario["age"] = "32"

	fmt.Println(usuario)


	slice := []int{0, 5, 10}

	slice = append(slice, 90, 110, 50)

	fmt.Println(slice)

}