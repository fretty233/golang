package main

import "fmt"

func main() {
	var name string

	name = "Fretty Silalahi"
	fmt.Println(name)

	name = "Fretty Mentari"
	fmt.Println(name)

	var word = "Silalahi"
	fmt.Println(word)

	var age int8 = 20
	fmt.Println(age)

	names := "fretty"
	fmt.Println(names)
	//:= deklarasi awal

	var (
		firstName = "Fretty"
		lastName  = "Silalahi"
	)
	fmt.Println(firstName + " " + lastName)

}
