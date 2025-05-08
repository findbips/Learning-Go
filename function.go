package main

import "fmt"

func cal(num1 int, num2 int, num3 int) {
	sum := num1 + num2 + num3

	fmt.Println(sum)
}

func main() {
	x := 20
	y := 30
	z := 50

	cal(x, y, z)
}