package main

import "fmt"

func v_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	name := "Golang"
	fmt.Println(name)

	fmt.Println()
}

func v_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y
	fmt.Println(a, x, y)
	fmt.Println()
}

func v_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, -15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error

	fmt.Println()
}

func v_exercise1() {
	fmt.Println("\n### Exercise 1 ###")
	var a int
	var b float64
	var c bool
	var d string
	fmt.Println(a, b, c, d)

	x := 20
	y := -15.5
	z := "Gopher!"
	fmt.Println(x, y, z)

	fmt.Println()
}
