package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Printf("%d", i)
		i = i + 1
	}
	fmt.Println()

	for j := 0; j < 3; j++ {
		fmt.Printf("%d", j)
	}
	fmt.Println()

	for i := range 3 {
		fmt.Printf("Range: %d\n", i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}
		fmt.Printf("%d", n)
	}
	fmt.Println()

}
