package main

import "fmt"

type user struct {
	name string
	age	 uint8
}

func (u user) save() {
	fmt.Printf("Saving data of User %s in the database\n", u.name)
}

func (u user) majorAge() bool {
	return u.age >= 18
}

func (u *user) userBirthday() {
	u.age++
}



func main() {
	user1 := user{"Victor", 25}
	user1.save()

	user1.userBirthday()
	user1.majorAge()
}