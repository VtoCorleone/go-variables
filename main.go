package main

import (
	"fmt"
)

func firstDeclaration() {
	var s, t string = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)
}

func secondDeclaration() {
	var (
		a string = "foo"
		i int    = 4
	)
	fmt.Println(a)
	fmt.Println(i)
}

func zeroValue() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func ifStatement() {
	var s string
	if s == "" {
		fmt.Println("s has not been assigned a value and is zero valued")
	}
}

func shortHand() {
	s := "Hello World"
	fmt.Println(s)
}

func differentDeclarations() {
	var s string = "Hello World"
	var a = "Hello World"
	var t string
	t = "Hello World"
	u := "Hello World"
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(t)
	fmt.Println(u)
}

func memoryLocation() {
	s := "Hello World"
	fmt.Println(&s)
}

func passByValue(x int) {
	fmt.Println(&x)
	return
}

func passByReference(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	return
}

func main() {
	firstDeclaration()
	secondDeclaration()
	zeroValue()
	ifStatement()
	shortHand()
	differentDeclarations()
	memoryLocation()

	i := 1
	fmt.Println(&i)
	passByValue(i)

	j := 1
	fmt.Println(&j)
	passByReference(&j)
}
