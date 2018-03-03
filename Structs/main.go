package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Nick", lastName: "Val"}
	fmt.Println(alex)
}
