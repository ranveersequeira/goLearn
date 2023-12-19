package main

import "fmt"

var num = 10

func main() {
	//if else only - change num value to play around;
	if num%2 == 0 {
		fmt.Println("we found the even number", num)
	} else {
		fmt.Println("how odd it is", num)
	}

	if num%2 == 0 {
		fmt.Println("Number is divisible by 2", num)
	}

	if num%2 == 0 && num%5 == 0 {
		fmt.Println("Number is divisible by 2 or 5 (change the num please!)", num)
	}

	if num == 10 {
		fmt.Println("for godsake change the num value", num)
	} else if num < 10 {
		fmt.Println("good you changed it to 1 digit", num)
	} else {
		fmt.Println("you're good at this keep up", num)
	}

}
