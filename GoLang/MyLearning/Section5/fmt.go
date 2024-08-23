package main

import (
	"fmt"
)

func fmt_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// Print each variable using a specific verb for its type
	fmt.Printf("x: %d, y: %f, z: %s, score: %#v\n", x, y, z, score)

	// Print the string value enclosed by double quotes ("Gophers")
	fmt.Printf("%q\n", z)

	// Print each variable using the same verb
	fmt.Printf("x: %v, y: %v, z: %v, score: %v\n", x, y, z, score)

	// Print the type of y and score
	fmt.Printf("x: %T, y: %T, z: %T, score: %T\n", x, y, z, score)

	fmt.Println()
}

func fmt_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	const x float64 = 1.422349587101
	fmt.Printf("%.4f", x)

	fmt.Println()
}

func fmt_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	// fmt.Printf("Shape: %q\n")
	fmt.Printf("Shape: %q\n", shape)
	// fmt.Printf("Circle's circumference with radius %d is %b\n", radius, circumference)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape

	fmt.Println()
}
