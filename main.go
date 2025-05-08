package main

import "fmt"

// func main() {
// 	age := 18

// 	if age > 18{
// 		fmt.Println("True")
// 	} else if age < 18 {
// 		fmt.Println("False")
// 	} else if age == 18 {
// 		fmt.Println("Equal")
// 	}

// }

// Operator
// > , >= , < , <= , == 
// and => &&
// or => ||
// not => !

func main() {
	age := 20
	sex := "male"

	if age == 20 || sex == "female" {
		fmt.Println("You're not elegible for this job.")
	} else {
		fmt.Println("Date not matched.")
	}

	isPretty := false

	if !isPretty {
		fmt.Println("Show something!")
	}
}