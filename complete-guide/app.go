package main

import "fmt"

func main() {
	hell_world()
	values()
	vars()
}

func hell_world() {
	fmt.Print("Hello, world!")
}

func values() {
	// Some basic values.
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func vars() {
	var a = "inital"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "Apple" // := operator is only accessable inside of functions
	fmt.Println(f)
}
