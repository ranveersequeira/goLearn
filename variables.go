package main

import "fmt"

func main() {
	//variable declared with var
	var firstName = "Ranveer"
	fmt.Println("This is my name -", firstName)

	//variable should have type or initialized with value so that go can infer its type like above
	var lastName, noName string = "Sequeira", "None"
	fmt.Println("Last Name and NoName -", lastName, noName)

	//we can also declare variable using this syntax(only works wraped in fn)
	favoriteFruit := "don't know yet"
	fmt.Println("What is your favorite fruit -", favoriteFruit)

}
