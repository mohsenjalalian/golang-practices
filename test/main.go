package main

import "fmt"

type person struct {
	firstname string
	lastname string
	contactinfo
}

type contactinfo struct {
	email string
	zipCode int
}

func main()  {
	jim := person {
		firstname: "jim",
		lastname: "party",
		contactinfo: contactinfo {
			email: "test",
			zipCode: 123,
		},
	}

	fmt.Println("")
	jim.update("gg")
	fmt.Println("")
	jim.print()
	fmt.Println("")
}

func (p person) update(newvalue string) {
	p.firstname = newvalue
	fmt.Printf("%+v", p)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}