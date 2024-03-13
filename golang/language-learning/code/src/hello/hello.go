package main

import "fmt"

const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "
const spanish = "Spanish"

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	if language == spanish {
// 		return spanishHello + name
// 	}
// 	return englishHello + name
// }

// 3.13 refactor code

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHello
	case "Spanish":
		prefix = spanishHello
	case "English":
		prefix = englishHello
	}
	return
}
func main() {
	fmt.Println(Hello("Bruce", "English"))
}
