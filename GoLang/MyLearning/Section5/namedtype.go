package main

import (
	"fmt"
)

type duration int

func n_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var hour duration
	fmt.Printf("%T : %d\n", hour, hour)

	hour = 3600
	fmt.Printf("%T : %d\n", hour, hour)

	fmt.Println()
}

func n_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))

	fmt.Println()
}

func n_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	type mile float64
	type kilometer float64
	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Println(kmBerlinToParis)
	fmt.Println(mileBerlinToParis * m2km)

	fmt.Println()
}
