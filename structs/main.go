package main

import "fmt"


type contactInfo  struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main(){

	// p1 := person{"name", "lastName"}
	// p2 := person{firstName: "name", lastName:  "lastName"}
	// var p3 person
	// p3.firstName = "name 3"
	// p3.lastName = "lastname 3"

	// fmt.Printf("%+v", p3)

	p4 := person{
		firstName: "Nombre",
		lastName: "last",
		contactInfo: contactInfo{
			email: "q@q.com",
			zipCode: 123000,
		},
	}
	// // Memory address of the variable
	// p4Pointer := &p4
	// fmt.Println(&p4)
	// // Update the value from memory
	// p4Pointer.updateName("Jose")
	// If the function recibir is declare as a pointer a shot cut can be use like this
	 p4.updateName("Jose")

	
	p4.print()
}

func (pointerPerson *person) updateName(newFirstName string){
	(*pointerPerson).firstName = newFirstName
}


func (p person) print(){
	fmt.Printf("%+v", p)
}