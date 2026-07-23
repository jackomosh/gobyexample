package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 9%2 != 0 && 7%2 != 0 {
		fmt.Println("Both 8 and 7 are Odd")
	}

	if num := 9; num < 0 {
		fmt.Println("Num is Negative")
	} else if num < 10 {
		fmt.Println("Num is 1 Digit")
	} else {
		fmt.Println("Num has multiple digits")
	}

}
