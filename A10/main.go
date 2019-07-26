package main

import "fmt"

func main()  {
	u := user{
		Name: "Helio",
		Age: 32,
		Live: true,
	}

	fmt.Println(u)

	u.Morrer()

	fmt.Println(u)

}

type user struct {
	Name string
	Age int
	Live bool
}

//Passando o ponteiro com o *
func (u *user) Morrer() {
	u.Live = false
}