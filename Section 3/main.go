package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	// selcuk := person{"Selcuk", "Candan"}
	// fmt.Println(selcuk)
	// fmt.Println(selcuk.firstname)
	// fmt.Println(selcuk.lastname)

	var john person
	john.firstname = "John"
	john.lastname = "Doe"
	fmt.Printf("%+v\n", john)

	selcuk := person{
		firstname: "Selcuk",
		lastname:  "Candan",
		contact: contactInfo{
			email:   "mehmetselcukcandan@icloud.com",
			zipCode: 11111,
		},
	}
	// fmt.Printf("%+v\n", selcuk)
	selcuk.print()
	selcuk.updateFirstname("Mehmet Selcuk")
	selcuk.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstname(newFirstname string) {
	p.firstname = newFirstname
}
