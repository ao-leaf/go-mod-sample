package main

import "fmt"

func Hello() {
	fmt.Printf("Hello")
}

func hello() {
	fmt.Printf("hello")
}

func main() {
	fmt.Printf("This is main function.")
	Hello()
	hello()
}