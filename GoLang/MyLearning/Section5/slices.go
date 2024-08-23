package main

import (
	"fmt"
	"os"
	"strconv"
)

func s_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	var s = []string{"a", "b", "c"}
	for idx, val := range s {
		fmt.Println(idx, "->", val)
	}

	fmt.Println()
}

func s_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	mySlice := []float64{1.2, 5.6}

	//mySlice[2] = 6
	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)

	a := 10
	mySlice[0] = float64(a)
	fmt.Println(mySlice)

	//mySlice[3] = 10.10
	mySlice = append(mySlice, 10.10)
	fmt.Println(mySlice)

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

	fmt.Println()
}

func s_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	nums := []float64{1.1, 2.2, 3.3}
	nums = append(nums, 10.1)
	fmt.Println(nums)

	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)

	n := []float64{11.11, 22.22}
	nums = append(nums, n...)
	fmt.Println(nums)

	fmt.Println()
}

func s_exercise4() {
	fmt.Println("\n### Exercise 4 ###")

	count := len(os.Args)
	if count < 3 || count > 11 {
		fmt.Printf("Got %d # of args, need between 2 to 10 only!!!", count)
	}

	sum := 0
	mult := 1
	for idx := 1; idx < len(os.Args); idx++ {
		num, _ := strconv.Atoi(os.Args[idx])
		sum += num
		mult *= num
	}
	fmt.Printf("Sum: %d, Product: %d\n", sum, mult)

	fmt.Println()
}

func s_exercise4_1() {
	fmt.Println("\n### Exercise 4.1 ###")

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	fmt.Println()
}

func s_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, val := range nums[1 : len(nums)-2] {
		fmt.Println(val)
		sum += val
	}
	fmt.Println("Sum:", sum)

	fmt.Println()
}

func s_exercise6() {
	fmt.Println("\n### Exercise 6 ###")

	friends := []string{"Marry", "John", "Paul", "Diana"}
	// var f []string
	var f = make([]string, len(friends))
	copy(f, friends)
	fmt.Println("Friends:", friends, "=> copy", f)

	fmt.Println()
}

func s_exercise7() {
	fmt.Println("\n### Exercise 7 ###")

	friends := []string{"Marry", "John", "Paul", "Diana"}
	f := []string{}
	f = append(f, friends...)
	fmt.Println(friends, "=>", f)
	f[1] = "Manish"
	fmt.Println(friends, "=>", f)

	// try
	f1 := friends[:]
	f1[1] = "Manish"
	fmt.Println(friends, "=>", f1)
	fmt.Println()
}

func s_exercise8() {
	fmt.Println("\n### Exercise 8 ###")

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := append(years[:3], years[len(years)-3:]...)

	// Another way
	// var newYears []int.
	// newYears = append(newYears, years[:3]...)
	// newYears = append(newYears, years[len(years)-3:]...)
	fmt.Println(newYears)

	fmt.Println()
}
