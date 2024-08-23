package main

import (
	"fmt"
)

func st_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	type person struct {
		name           string
		age            int
		favoriteColors []string
	}

	me := person{"Manish", 40, []string{"Yellow", "Black"}}
	you := person{name: "Shweta", age: 39, favoriteColors: []string{"a", "b", "c"}}
	fmt.Println(me)
	fmt.Println(you)

	fmt.Println()
}

func st_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	type person struct {
		name           string
		age            int
		favoriteColors []string
	}

	me := person{"Manish", 40, []string{"Yellow", "Black"}}
	you := person{name: "Shweta", age: 39, favoriteColors: []string{"a", "b", "c"}}
	fmt.Println(me)
	fmt.Println(you)

	me.name = "Andrei"
	fmt.Printf("%#v\n", me)

	colors := you.favoriteColors
	fmt.Printf("%T -- %v\n", colors, colors)

	you.favoriteColors = append(you.favoriteColors, "Pink")
	fmt.Println(you)
	fmt.Printf("%+v\n", you)
	fmt.Printf("%+v\n", me)

	fmt.Println()
}

func st_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	type person struct {
		name           string
		age            int
		favoriteColors []string
	}

	me := person{"Manish", 40, []string{"Yellow", "Black"}}
	you := person{name: "Shweta", age: 39, favoriteColors: []string{"a", "b", "c"}}
	fmt.Println(me)
	fmt.Println(you)

	for i, v := range me.favoriteColors {
		fmt.Println(i, "--", v)
	}

	fmt.Println()
}

func st_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	type grades struct {
		grade, course string
	}

	type person struct {
		name           string
		age            int
		favoriteColors []string
		grade          grades
	}

	me := person{"Manish", 40, []string{"Yellow", "Black"}, grades{"B", "MCA"}}
	you := person{name: "Shweta", age: 39, favoriteColors: []string{"a", "b", "c"}, grade: grades{"A", "BHMS"}}
	fmt.Println(me)
	fmt.Println(you)

	me.grade.grade = "98"
	me.grade.course = "Golang"

	fmt.Printf("%+v\n", you)
	fmt.Printf("%+v\n", me)

	fmt.Println()
}
