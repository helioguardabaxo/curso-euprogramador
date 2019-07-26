package main

import "fmt"

func main()  {
	sl := make([]int, 2, 3)

	sl[0] = 10
	sl[1] = 13

	sl2 := sl[:]

	sl = append(sl, 1, 60)

	sl[0] = 200

	fmt.Println(sl)
	fmt.Println(sl2)
}