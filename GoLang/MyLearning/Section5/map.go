package main

import (
	"fmt"
)

func m_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var m1 map[int]int
	fmt.Printf("%T -- %#v\n", m1, m1)

	var m2 = map[int]string{
		1: "Manish",
		2: "Shweta",
	}
	m2[10] = "Abba"
	fmt.Println(m2)

	fmt.Println(m2[2])

	fmt.Println(m2[4])

	fmt.Println()
}

func m_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	var m1 = map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m2 == m3)
	fmt.Println(m1, m2, m3)

	fmt.Println()
}

func m_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	m := map[int]bool{1: true, 2: false, 3: false}
	for idx, val := range m {
		fmt.Println(idx, "==", val)
	}

	delete(m, 2)

	for idx, val := range m {
		fmt.Println(idx, "==", val)
	}

	fmt.Println()
}
