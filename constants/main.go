package main

import (
	"fmt"
	"math"
)

const n = 500000000

func main() {
	const name string = "Jack"
	// name := "Omondi"
	fmt.Println(name)

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}