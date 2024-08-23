package main

import "fmt"

func main() {
	var name string = "Manish"
	dob := "Sep1982"
	var city = "Jaipur"
	fmt.Println(name, dob, city)

	var m = 345
	fmt.Printf("%#v\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%d\n", m)

	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c) // => a is 10, b is 15.500000, c is Gophers
	fmt.Printf("%q\n", c)                               // => "Gophers"
	fmt.Printf("%v\n", grades)                          // => [10 20 30]
	fmt.Printf("%#v\n", grades)                         // => b is of type float64 and grades is of type []int
	fmt.Println("====================")

	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128
	fmt.Printf("%c and %c\n", 100, 51011)    // =>  d and 읃  (runes for code points 101 and 51011)

	const pi float64 = 3.14159265359
	fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points

	const (
		min1 = -500
		max1 //gets its type and value form the previous constant. It's 500
		max2 //in a grouped constants, a constant repeats the previous one -> 500
	)
	fmt.Println(min1, max1, max2)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	const (
		North = iota //by default 0
		East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)
	fmt.Println(North, West)

	var i3 int64 = -324_567345   // underscores are used to write large numbers for a better readability
	fmt.Printf("%T\n", i3)       // => int64
	fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)

	//rune type
	var r rune = 'f'
	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
	fmt.Printf("%c\n", r) // => f
	fmt.Printf("%v\n", r) // => 102
	fmt.Printf("%d\n", r) // => 102

	//slice type
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // => []string
	fmt.Println(append(cities, "Jaipur"))
	fmt.Printf("%v\n", cities) // => []string

	//struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you) // => main.Person

	//pointer type
	var x int = 2
	ptr := &x                                                 // pointer to int
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168

	//function type
	fmt.Printf("%T\n", f) // => func()

}

func f() {
}
