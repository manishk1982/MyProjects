package main

import (
	"fmt"
	"unicode/utf8"
)

func str_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var name string = "Manish"
	country := "USA"
	fmt.Println("Your name:", name)
	fmt.Println("Country:", country)

	fmt.Println("a) He says: \"Hello\"")
	fmt.Println("b) C:\\Users\\a.txt")

	fmt.Println(`a) He says: "Hello"`)
	fmt.Println(`b) C:\Users\a.txt`)

	fmt.Println()
}

func str_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	var r rune = 'ă'
	fmt.Printf("%T\n", r)
	fmt.Printf("%c\n", r)

	s1, s2 := "ma", "m"
	s3 := s1 + s2 + string(r)
	fmt.Println(s3)

	fmt.Println()
}

func str_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v  -  ", s1[i])
	}

	fmt.Println()
	for idx, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", idx, r)
	}

	fmt.Println()
}

func str_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	fmt.Println()
}

func str_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	s := "你好 Go!"
	b := []byte(s)
	fmt.Println(b)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%d\n", b)
	fmt.Printf("%c\n", b)

	for i, r := range b {
		fmt.Printf("Index: %d -- valie: %d\n", i, r)
	}

	fmt.Println()
}

func str_exercise6() {
	fmt.Println("\n### Exercise 6 ###")

	s := "你好 Go!"
	r := []rune(s)
	fmt.Println(r)
	fmt.Printf("%#v\n", r)

	for i, r1 := range r {
		fmt.Printf("Index: %d -- valie: %c\n", i, r1)
	}

	fmt.Println()
}
