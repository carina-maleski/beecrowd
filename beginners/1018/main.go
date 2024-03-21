package main

/*
In this problem you have to read an integer value and calculate the smallest possible number of banknotes in which the value may be decomposed. The possible banknotes are 100, 50, 20, 10, 5, 2 and 1. Print the read value and the list of banknotes.

Input
The input file contains an integer value N (0 < N < 1000000).

Output
Print the read number and the minimum quantity of each necessary banknotes in Portuguese language, as the given example. Do not forget to print the end of line after each line, otherwise you will receive “Presentation Error”.
*/

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