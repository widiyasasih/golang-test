package main

import "fmt"

func main() {
	fmt.Println("Hello! Welcome to quiz game!\nPlease introduce yourself...\n")

	var name string
	var age int
	var answer1 string
	var answer2 string
	score := 0

	fmt.Printf("Your name: ")
	fmt.Scan(&name)
	fmt.Printf("Your age: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Println("\nYou need to be 12 yo or older. See you another time!")
	} else {
		fmt.Printf("\n+++ Great %v let's begin! +++", name)
	}

	fmt.Printf("\nWhat the most powerful language between Python and C for machine learning? ")
	fmt.Scan(&answer1)
	if answer1 == "Python" || answer1 == "python" {
		fmt.Println("=> Correct!")
		score += 1
	} else {
		fmt.Println("=> Incorrect!")
	}

	fmt.Printf("\nIf only need developer for the testing, what should be used between beta testing or alpha testing? ")
	fmt.Scan(&answer2)
	if answer2 == "alpha" || answer2 == "alpha testing" {
		fmt.Println("=> Correct!")
		score += 1
	} else {
		fmt.Println("=> Incorrect!")
	}

	fmt.Println("\nYour score: ", score)
}
