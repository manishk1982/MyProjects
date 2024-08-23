package main

import (
	"fmt"
)

func a_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var cities [2]string
	fmt.Printf("%#v\n", cities)

	var grades = [3]float64{
		3.4,
		4.5,
	}
	fmt.Printf("%#v\n", grades)

	var taskDone = [...]bool{}
	fmt.Printf("%#v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])

	}

	for _, val := range grades {
		fmt.Println(val)
	}

	fmt.Println()
}

func a_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	nums := [...]int{30, -1, -6, 90, -6}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i]%2 == 0 {
			fmt.Println(nums[i])
			count++
		}
	}
	fmt.Println("No. of positive even numbers in nums: ", count)

	fmt.Println()
}

func a_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6
	// myArray[2] = float64(6)

	a := 10
	myArray[0] = float64(a)

	// myArray[3] = 10.10
	myArray[2] = 10.10

	fmt.Println(myArray)

	fmt.Println()
}
