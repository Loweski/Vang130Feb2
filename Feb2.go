// Vang130Feb2 project Feb2.go
// Vangtest project Feb2.go
package main

import (
	"fmt"
)

type Person struct {
	name     string
	position string
	course   string
}

func main() {
	var person1 = Person{}
	var person2 = Person{}
	var exam1, exam2 = 80, 90
	var avgscore int
	avgscore = (exam1 + exam2) / 2
	message := "Person:"
	person1.name = "Leng"
	person1.position = "Student"
	person1.course = "Csci 130"

	person2.name = "Todd"
	person2.position = "Professor"
	person2.course = "Csci 130"

	fmt.Println(message, person1.name, person1.position, person1.course)
	fmt.Print("Score:")
	fmt.Println(avgscore)
	var greeting *string = &message

	*greeting = "Hello"
	fmt.Println(message, person2.name, "!")
}
