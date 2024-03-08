package main

import "fmt"

const englishHello = "Hello, "

func Hello(name string) string {
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Bruce"))
}
