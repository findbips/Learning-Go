package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to the web application!")
}
func getUserName() string {
	var name string
	fmt.Println("Enter your full name : ")
	fmt.Scanln(&name)

	return name
}
func getTwoNumbers()(int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number : ")
	fmt.Scanln(&num1) // & => ampersand
	fmt.Println("Enter second number : ")
	fmt.Scanln(&num2)

	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func displayResults(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)

}
func closingMessage(name string) {
	fmt.Println("Thankyou for using the application,", name)
	fmt.Println("Have a good day!")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	displayResults(name, sum)
	closingMessage(name)
}