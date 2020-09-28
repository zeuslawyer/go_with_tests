package main

import "fmt"

const ENG = "English"
const SPANISH = "Spanish"
const FRENCH = "French"

func Hello(name string, language string) (greeting string) {
	if name == "" {
		name = "World"
	}

	suffix := name + "!"
	greeting = makePrefix(language) + suffix
	return greeting
}

func makePrefix(language string) (prefix string) {
	switch language {
	case FRENCH:
		prefix = "Bonjour, "
	case SPANISH:
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}
	return prefix
}

func main() {
	fmt.Println("Hello, World! I am learning me some Golang!")
}
