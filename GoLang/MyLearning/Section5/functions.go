package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f_exercise1() {
	fmt.Println("\n### Exercise 1 ###")

	fmt.Println(cube(3))
	fmt.Println()
}
func cube(f float64) float64 {
	return f * f * f
}

func f_exercise2() {
	fmt.Println("\n### Exercise 2 ###")

	fmt.Println(f1(4))

	fmt.Println()
}
func f1(n uint) (uint, uint) {
	var sum uint = 0
	var fact uint = 1
	var i uint = 1
	for i <= n {
		sum += i
		fact *= i
		i++
	}
	return fact, sum

}

func f_exercise3() {
	fmt.Println("\n### Exercise 3 ###")

	fmt.Println(myFunc("5"))

	fmt.Println()
}
func myFunc(s string) int {
	n, _ := strconv.Atoi(s)
	nn, _ := strconv.Atoi(s + s)
	nnn, _ := strconv.Atoi(s + s + s)
	return n + nn + nnn
}

func f_exercise4() {
	fmt.Println("\n### Exercise 4 ###")
	fmt.Println(sum(3, 4, 5, 6))

	fmt.Println()
}
func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func f_exercise5() {
	fmt.Println("\n### Exercise 5 ###")

	fmt.Println(nakedreturn(3, 4, 5, 6))

	fmt.Println()
}
func nakedreturn(n ...int) (sum int) {
	sum = 0
	for _, v := range n {
		sum += v
	}
	return
}

func f_exercise6() {
	fmt.Println("\n### Exercise 6 ###")

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	fmt.Println()
}
func searchItem(str_slice []string, s string) bool {
	flag := false
	for _, v := range str_slice {
		if s == v {
			flag = true
		}
	}
	return flag
}

func f_exercise7() {
	fmt.Println("\n### Exercise 7 ###")

	animals := []string{"Lion", "tiger", "bear"}
	result := searchItemCase(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItemCase(animals, "lion")
	fmt.Println(result) // => true

	fmt.Println()
}
func searchItemCase(str_slice []string, s string) bool {
	flag := false
	for _, v := range str_slice {
		if strings.EqualFold(s, v) {
			flag = true
		}
	}
	return flag
}

func f_exercise8() {
	fmt.Println("\n### Exercise 8 ###")

	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")

	fmt.Println()
}
func print(msg string) {
	fmt.Println(msg)
}

func f_exercise9() {
	fmt.Println("\n### Exercise 9 ###")

	pp(6)
	var fx = pp
	fmt.Printf("%T\n", fx)

	fx(8)

	fmt.Println()
}
func pp(i int) {
	fmt.Println(i)
}
