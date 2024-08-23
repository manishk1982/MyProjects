package main

import (
	"fmt"
)

func fc_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i, "is divisible by 7")
		}
	}

	fmt.Println()
}

func fc_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i, "is divisible by 7")
	}

	fmt.Println()
}

func fc_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i, "is divisible by 7")
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println()
}

func fc_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i, "is divisible by 7 & 5")
		}
	}
	fmt.Println()
}

func fc_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	curYear := 2023
	for year := 1982; year <= curYear; {
		fmt.Println(year)
		year++
	}
}

func fc_exercise6() {
	fmt.Println("\n### Exercise 6 ###")

	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age < 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
	fmt.Println()
}
