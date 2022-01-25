package main

import "fmt"

func main() {
	var nilai32 int32 = 1000000000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name string = "Fretty Silalahi"
	var a byte = name[7]
	var astring string = string(a)

	fmt.Println("isi name", name)
	fmt.Println("isi astring", astring)
	fmt.Println("nilai a", a)
}
