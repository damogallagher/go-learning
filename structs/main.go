package main

import "fmt"

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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@example.com",
			zipCode: 94000,
		},
	}

	// &: gives the memory address of the value this variable is pointing at
	//alexPointer := &alex
	//alexPointer.updateName("alexander")
	//pointer shortcut
	alex.updateName("alexander")
	alex.print()
}

// *: gives the value the memory address is pointing at
// * on a type (person) means we are working with a pointer to a pointerToPerson
// * on a value (*pointerToPerson) is an operator - means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v\n\n", p)
}
