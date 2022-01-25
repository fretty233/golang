package main

import "fmt"

func main() {
	type NoKtp string
	type married bool

	var ktpku NoKtp = "1233123123213"
	var marriedstatus married = false

	fmt.Println(ktpku)
	fmt.Println(marriedstatus)
}
