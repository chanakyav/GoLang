package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	nick := person{
		firstName: "Chanakya",
		lastName:  "Val",
		contactInfo: contactInfo{
			email:   "cvalluri@email.com",
			zipCode: 95129,
		},
	}

	nick.updateName("Nick")
	nick.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
