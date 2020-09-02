package main

import "fmt"

type contactList struct {
	id          int
	name        string
	phoneNumber string
	email       string
	created_at  string
}

func (c contactList) display() {
	fmt.Println(
		c.id,
		c.name,
		c.phoneNumber,
		c.email,
		c.created_at,
	)
}

func main() {
	contact := contactList{
		id:          1,
		name:        "John",
		phoneNumber: "998901234567",
		email:       "john@doe.com",
		created_at:  "2020/02/20"}
	contact.display()
}
