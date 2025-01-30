package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex@gamil.com"
	alex.contact.zipcode = 12345

	alex.print()
	alex.update_name(("Jim"))
	alex.print()
}

func (p person) print() {
	fmt.Print(p)
}

func (pointerToPerson *person) update_name(new_name string) {
	(*pointerToPerson).firstName = new_name
}
