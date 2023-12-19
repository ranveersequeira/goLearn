package main

import "fmt"

func main() {
	//for loop is only way to loop in go,
	//for loop with limit
	i := 0
	for i < 6 {
		fmt.Println("Index", i)
		i += 1
	}

	//traditional for loop
	for j := 0; j < 10; j++ {
		fmt.Println("trad loop", j)
	}

	// clone of while loop
	for {
		fmt.Println("Only (god) return or break can save me")
		//don't Uncomment it would break your system
		break // return

	}

	//conditional loop
	for n := 1; n < 10; n++ {
		if n%2 == 0 {
			console.log("gotcha, comment it out")
			fmt.Println("good number", n)
		}
		continue
	}
}
