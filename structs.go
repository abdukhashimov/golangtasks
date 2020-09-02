package main

import "fmt"

type tasks struct {
	id   int
	name string
	done bool
}

func (t tasks) display() {
	fmt.Println("ID: ", t.id, "Name: ", t.name, "Done: ", t.done)
}

func main() {

}
