package main

import "fmt"

func main()  {
	var adm Admin

	adm.ranking = 1

	adm.name = "Helio"
	fmt.Println(adm.name)

	adm.Login()

	fmt.Println(adm.online)

	adm.Logout()

	fmt.Println(adm.online)

}

type User struct {
	name string
	online bool
}

func (u *User) Login()  {
	u.online = true
}

func (u *User) Logout()  {
	u.online = false
}

/*O tipo Admin extende o tipo User, herdando os atributos
name e online*/
type Admin struct {
	User
	ranking int
}