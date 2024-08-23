package main

import (
	"fmt"
)

func c_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	const daysWeek int = 7

	const lightSpeed int64 = 299792458

	const pi float64 = 3.14159
	fmt.Println(daysWeek, lightSpeed, pi)

	fmt.Println()
}

func c_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
	fmt.Println(daysWeek, lightSpeed, pi)

	fmt.Println()
}

func c_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	// Declare secPerDay constant and initialize it to the number of seconds in a day
	const secPerDay = 60 * 60 * 24

	//  2. Declare daysYear constant and initialize it to 365
	const daysYear = 365

	//  3. Use fmt.Printf() to print out the total number of seconds in a year.
	fmt.Printf("Total number of seconds in a year: %d\n", secPerDay*daysYear)

	fmt.Println()
}

func c_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	const x int = 10

	// declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8}
	m := []int{1, 3, 4, 5, 6, 8}
	_ = m

	fmt.Println()
}

func c_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	// const a int = 7
	const a = 7
	const b float64 = 5.6
	const c = a * b
	fmt.Println(a, b, c)

	x := 8
	_ = x
	// const xc int = x

	// const noIPv6 = math.Pow(2, 128)

	fmt.Println()
}

func c_exercise6() {
	fmt.Println("\n### Exercise 6 ###")

	// const (
	// 	_   = iota
	// 	_   = iota
	// 	_   = iota
	// 	_   = iota
	// 	_   = iota
	// 	_   = iota
	// 	jun = iota
	// 	jul = iota
	// 	aug = iota
	// )
	const (
		jun = iota + 6
		jul //automatically incremented by one
		aug
	)
	fmt.Println(jun, jul, aug)

	fmt.Println()
}
