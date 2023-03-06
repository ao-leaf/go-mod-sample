package main

import "fmt"

func Hello() {
	fmt.Printf("Hello")
}

func hello() {
	fmt.Printf("hello")
}

func SayHello() {
	Hello()
	hello()
}
