package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	//operasi matematika
	var c = a + b
	var d = a * b
	var e = a / b
	var f = a % b
	var g = a - b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("penjumalahan = ", c)
	fmt.Println("perkalian = ", d)
	fmt.Println("pembagian = ", e)
	fmt.Println("modulo = ", f)
	fmt.Println("pengurangan = ", g)

	//augmented assigments
	var i = 10
	var j = 10
	i += 10
	fmt.Println("i+", i)

	j -= 10
	fmt.Println("j-", j)
	fmt.Println("i-", i)

}
