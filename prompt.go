package main

import "fmt"

func main() {
	var firstname string
	fmt.Print("whats your name\n")
	fmt.Scan(&firstname)
	fmt.Print("whats ur last name\n")
	var lastname string
	fmt.Scan(&lastname)
	fmt.Print("how old r u\n")
	var age int
	fmt.Scan(&age)
	// print the name and age
	fmt.Printf("hello, %s %s! you are %d years old\n", firstname, lastname, age)
}
