package main

import (
	"fmt"
	"strconv"
)

func o_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var i int = 3
	var f float64 = 3.2

	var fi = float64(i)
	fmt.Printf("type: %T, value: %f\n", fi, fi)

	var intf = int(f)
	fmt.Printf("type: %T, value: %d\n", intf, intf)

	fmt.Println()
}

func o_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// i to string (int to string)
	itos := strconv.Itoa(i)
	fmt.Printf("type: %T, value: %s\n", itos, itos)

	// s2 to int (string to int)
	s2toi, err := strconv.Atoi(s2)
	_ = err
	fmt.Printf("type: %T, value: %v\n", s2toi, s2toi)

	// f to string (float64 to string)
	f2s := fmt.Sprintf("%f", f)
	fmt.Printf("type: %T, value: %s\n", f2s, f2s)

	// s1 to float64  (string to float64)
	s12f, err := strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Printf("type: %T, value: %f\n", s12f, s12f)

	fmt.Println()
}

func o_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	x, y := 4, 5.1

	z := x * int(y)
	fmt.Println(z)

	z1 := float64(x) * y
	fmt.Println(z1)

	a := 5
	b := 6.2 * float32(a)
	fmt.Println(b)

	fmt.Println()
}

func o_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	// distance := 149.6
	// speed := 299_792_458
	// time1 := distance / float64(speed)
	// fmt.Println(time1)

	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speed = 299_792_458 // speed of light in m / s
	)
	const time = distance / speed // time in seconds
	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)

	fmt.Println()
}

func o_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z || int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1*5 && int(z) == 0
	fmt.Println(result2)

	fmt.Println()
}
