package main

import "fmt"

func cal(num1 int, num2 int, num3 int) int {
	sum := num1 + num2 + num3

	return sum
}

func getNumbers(num1 int, num2 int, num3 int)(int, int) {

	sum := num3 + num2 - num1
	mul := num1 * num2 / num3

	return sum, mul
}

func printSomething() {
	fmt.Println("Education must be free!")
}

func sayHello( name string) {
	fmt.Println("Welcome to goland course, ", name)
}


func main() {
	x := 20
	y := 30
	z := 50

	s, m := getNumbers(x, y, z)

	fmt.Println(s)
	fmt.Println(m)
	printSomething()
	sayHello("Biplob")
}