package main

import "fmt"

func main() {
	var input uint
	fmt.Scanln(&input)
	fmt.Println(input)

	var n100 uint = input/100
	input %= 100

	var n50 uint = input/50
	input %= 50

	var n20 uint = input/20
	input %= 20
	
	var n10 uint = input/10
	input %= 10
	
	var n5 uint = input/5
	input %= 5

	var n2 uint = input/2
	input %= 2
	
	var n1 uint = input

	fmt.Println(n100, "nota(s) de R$ 100,00")
	fmt.Println(n50, "nota(s) de R$ 50,00")
	fmt.Println(n20, "nota(s) de R$ 20,00")
	fmt.Println(n10, "nota(s) de R$ 10,00")
	fmt.Println(n5, "nota(s) de R$ 5,00")
	fmt.Println(n2, "nota(s) de R$ 2,00")
	fmt.Println(n1, "nota(s) de R$ 1,00")	
}